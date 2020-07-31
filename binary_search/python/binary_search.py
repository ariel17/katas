def binary_search(target, arr, init_pos=0):
    if len(arr) == 0:
        return -1

    if len(arr) == 1:
        if arr[0] == target:
            return init_pos
        return -1

    m = len(arr) // 2
    if arr[m] == target:
        return m+init_pos

    if arr[m] > target:
        return binary_search(target, arr[:m], 0)

    return binary_search(target, arr[m:], m)
