import { createStore } from 'vuex'

/**
 * 解析 AI 回复文本，拆分成：
 * - 纯文本（去除所有代码块后的文本）
 * - 代码块列表（包含语言和代码）
 * - Mermaid 代码块列表（过滤出语言为 mermaid 的代码块）
 * @param {string} text AI 回复文本
 * @returns {object} { pureText, codeBlocks, mermaidBlocks }
 */
function parseAIReply(text) {
  const codeBlockRegex = /```(\w+)?\n([\s\S]*?)```/g

  let lastIndex = 0
  let match
  const codeBlocks = []
  const mermaidBlocks = []
  const textParts = []

  while ((match = codeBlockRegex.exec(text)) !== null) {
    // 代码块前的纯文本部分
    if (match.index > lastIndex) {
      textParts.push(text.slice(lastIndex, match.index))
    }

    const lang = (match[1] || '').toLowerCase()
    const code = match[2].trim()

    codeBlocks.push({ lang, code })

    if (lang === 'mermaid') {
      mermaidBlocks.push(code)
    }

    lastIndex = codeBlockRegex.lastIndex
  }

  // 最后剩余纯文本
  if (lastIndex < text.length) {
    textParts.push(text.slice(lastIndex))
  }

  const pureText = textParts.join('\n').trim()

  return {
    pureText,
    codeBlocks,
    mermaidBlocks
  }
}

export default createStore({
  state: {
    chatHistory: [],       // { text: string, isUser: boolean }
    codeBlocks: [],        // [{ lang: string, code: string }]
    mindMap: [],           // [mermaidCode1, mermaidCode2, ...]
    settings: {
      modelApi: 'gpt-4',
      historyMessageCount: 10
    },
    knowledgeBase: {
      documents: [],       // 知识库文档列表
      isUploading: false,
      lastError: null
    },
    notification: {
      show: false,
      type: 'info',
      title: '',
      message: ''
    },
    conversationId: ''     // 会话ID
  },
  mutations: {
    ADD_MESSAGE(state, message) {
      state.chatHistory.push(message)
    },
    SET_CODE_BLOCKS(state, blocks) {
      state.codeBlocks = blocks
    },
    SET_MIND_MAP(state, mindMapArray) {
      state.mindMap = mindMapArray
    },
    UPDATE_MODEL_API(state, api) {
      state.settings.modelApi = api
    },
    UPDATE_HISTORY_COUNT(state, count) {
      state.settings.historyMessageCount = count
    },
    SET_KNOWLEDGE_DOCUMENTS(state, docs) {
      state.knowledgeBase.documents = docs
    },
    SET_UPLOADING_STATUS(state, status) {
      state.knowledgeBase.isUploading = status
    },
    SET_UPLOAD_ERROR(state, error) {
      state.knowledgeBase.lastError = error
    },
    SET_NOTIFICATION(state, notification) {
      state.notification = {
        ...state.notification,
        ...notification,
        show: true
      }
    },
    CLEAR_NOTIFICATION(state) {
      state.notification.show = false
    },
    SET_CONVERSATION_ID(state, id) {
      state.conversationId = id
    }
  },
  actions: {
    initConversationId({ commit }) {
      const savedId = localStorage.getItem('conversationId')
      if (savedId) {
        commit('SET_CONVERSATION_ID', savedId)
      }

    },

    async sendMessage({ commit, state }, message) {
      commit('ADD_MESSAGE', { text: message, isUser: true });

      try {
        const response = await fetch('http://localhost/v1/chat-messages', {
          method: 'POST',
          headers: {
            'Authorization': 'Bearer app-zCqAeaTjjb5uE4URqpamQDln',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            inputs: {},
            query: message,
            response_mode: 'blocking',
            conversation_id: state.conversationId || '',
            user: '111',
          })

        })

        if (!response.ok) throw new Error(`请求失败，状态码: ${response.status}`)

        const data = await response.json()
        const aiReply = data.answer || 'AI 未返回有效内容'

        if (data.conversation_id) {
          commit('SET_CONVERSATION_ID', data.conversation_id)
          localStorage.setItem('conversationId', data.conversation_id)
        }

        const { pureText, codeBlocks, mermaidBlocks } = parseAIReply(aiReply)

        if (pureText) {
          commit('ADD_MESSAGE', { text: pureText, isUser: false })
        }

        commit('SET_CODE_BLOCKS', codeBlocks)
        commit('SET_MIND_MAP', mermaidBlocks)

      } catch (error) {
        console.error('调用 AI 接口出错:', error)
        commit('ADD_MESSAGE', { text: '抱歉，AI 服务暂时不可用。', isUser: false })
        commit('SET_CODE_BLOCKS', [])
        commit('SET_MIND_MAP', [])
        throw error
      }
    },

    updateModelApi({ commit }, api) {
      commit('UPDATE_MODEL_API', api)
    },

    updateHistoryCount({ commit }, count) {
      commit('UPDATE_HISTORY_COUNT', count)
    },

    async fetchKnowledgeDocuments({ commit }, { isInitialLoad = false } = {}) {
      commit('SET_UPLOAD_ERROR', null)

      // 只有在非初始加载时才设置上传状态为true
      if (!isInitialLoad) {
        commit('SET_UPLOADING_STATUS', true)
      }

      try {
        const datasetId = "2750d917-3f80-47e6-ab14-01f239ebcf27"
        const apiKey = "dataset-3qH43ExwpT2qaYylvDNzIiZW"

        const res = await fetch(`http://localhost/v1/datasets/${datasetId}/documents`, {
          headers: {
            Authorization: `Bearer ${apiKey}`,
          },
        })

        if (!res.ok) {
          const errData = await res.json()
          throw new Error(errData.message || `请求失败，状态码${res.status}`)
        }

        const data = await res.json()
        commit('SET_KNOWLEDGE_DOCUMENTS', data.data || [])
      } catch (error) {
        commit('SET_UPLOAD_ERROR', error.message || '获取文档列表失败')
        commit('SET_KNOWLEDGE_DOCUMENTS', [])
      } finally {
        // 如果是初始加载过程中设置了上传状态，确保也重置它
        commit('SET_UPLOADING_STATUS', false)
      }
    },


    async uploadKnowledgeFile({ commit, dispatch }, { file, datasetId, apiKey }) {
      commit('SET_UPLOADING_STATUS', true)
      commit('SET_UPLOAD_ERROR', null)

      const dataObj = {
        indexing_technique: "high_quality",
        process_rule: {
          rules: {
            pre_processing_rules: [
              { id: "remove_extra_spaces", enabled: true },
              { id: "remove_urls_emails", enabled: true },
            ],
            segmentation: { separator: "###", max_tokens: 500 },
          },
          mode: "custom",
        },
      }

      const formData = new FormData()
      formData.append("data", JSON.stringify(dataObj))
      formData.append("file", file)

      try {
        const res = await fetch(
            `http://localhost/v1/datasets/${datasetId}/document/create-by-file`,
            {
              method: "POST",
              headers: {
                Authorization: `Bearer ${apiKey}`,
                // 不要手动设置 Content-Type
              },
              body: formData,
            }
        )

        if (!res.ok) {
          const errData = await res.json()
          throw new Error(errData.message || `上传失败，状态码${res.status}`)
        }

        const resData = await res.json()
        // 上传成功后刷新文档列表
        await dispatch('fetchKnowledgeDocuments')
        return resData
      } catch (error) {
        commit('SET_UPLOAD_ERROR', error.message || '上传失败')
        throw error
      } finally {
        commit('SET_UPLOADING_STATUS', false)
      }
    },

    showNotification({ commit }, { type, title, message }) {
      commit('SET_NOTIFICATION', { type, title, message })

      setTimeout(() => {
        commit('CLEAR_NOTIFICATION')
      }, 3000)
    }
  }
})
