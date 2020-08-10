def binary_search(target, arr):
    larr = len(arr)
    if larr == 0:
        return -1
    m = larr // 2
    left, found, right = arr[:m], arr[m] == target, arr[m+1:]
    if found:
        return m
    lfound, rfound = binary_search(target, left), binary_search(target, right)
    if lfound != -1:
        return lfound
    if rfound != -1:
        return m+1+rfound
    return -1
