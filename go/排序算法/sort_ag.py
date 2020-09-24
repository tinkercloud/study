
# i j
# 5 3 1 6 7 9

# 冒泡排序算法
def bubbleSort(alist):
    loop=0
    
    hasChange=True
    # 循环数组长度,作为第一层外循环
    for i in range(len(alist)):
        # 如果没有产生任何交换, 说明无需交换了,直接可以跳过
        if hasChange:
            hasChange=False
            # 内层循环,内层循环达到最后一个数组值
            for j in range(len(alist)-1):
                loop+=1
                if alist[j]>alist[j+1]:
                    tmp = alist[j]
                    alist[j]=alist[j+1]
                    alist[j+1]=tmp
                    hasChange=True
            
    print(loop)
    return alist
# 冒泡排序
res =bubbleSort([5,3,1,7,6,2,8,12,4,18,13,10])
print(res)

################################################################
