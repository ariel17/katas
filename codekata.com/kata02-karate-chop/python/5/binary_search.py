def binary_search(target, arr):
    found = False
    completed = False
    base = 0
    pos = -1
    while not found and not completed:
        m = len(arr) // 2
        if len(arr) == 0:
            completed = True
        elif arr[m] == target:
            found = True
            pos = base + m
        elif arr[m] > target:
            arr = arr[:m]
        else:
            base = m+1
            arr = arr[m+1:]
    return pos
