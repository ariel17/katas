def binary_search(target, arr):
    pos = -1
    base = 0
    found = False
    while len(arr) != 0 and not found:
        m = len(arr) // 2
        if arr[m] == target:
            pos = base + m
            found = True
        elif arr[m] > target:
            arr = arr[:m]
        else:
            base = m+1
            arr = arr[m+1:]
    return pos
