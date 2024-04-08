def reverse_array(arr):
    l=len(arr)
    for i in range(l//2):
        arr[i],arr[l-i-1]=arr[l-i-1],arr[i]
    return arr

def rotateLeft(arr, k):
    l = len(arr)
    k = k % l
    arr[:k] = reverse_array(arr[:k])
    arr[k:] = reverse_array(arr[k:])
    arr = reverse_array(arr)
    return arr

def rotateRight(arr, k):
    l = len(arr)
    k = k % l
    arr[:l-k] = reverse_array(arr[:l-k])
    arr[l-k:] = reverse_array(arr[l-k:])
    arr = reverse_array(arr)
    return arr

arr = [1, 2, 3, 4, 5, 6, 7]
rotateLeft(arr, 2)
print(arr)

arr = [1, 2, 3, 4, 5, 6, 7]
rotateRight(arr, 2)
print(arr)