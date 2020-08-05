def binary_search(target, arr):
    for (pos, value) in enumerate(arr):
        if value == target:
            return pos
    return -1
