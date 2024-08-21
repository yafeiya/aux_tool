import sys


def main():
    # 获取传入的参数
    args = sys.argv[1]  # 忽略第一个参数（脚本名）
    args2 = sys.argv[2]  # 忽略第一个参数（脚本名）
    # 打印参数
    print("Received arguments:", args,args2)
    
if __name__ == "__main__":
    main()