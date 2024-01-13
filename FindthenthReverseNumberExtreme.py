def find_reverse_number(n):
    p, sum_val = 18, 1

    while sum_val + p <= n - 1:
        sum_val += p
        p *= 10

    r = n - sum_val - 1
    p_div_2 = p // 2
    s1 = str(r % p_div_2 + p // 18)
    
    s2 = s1[::-1][r < p_div_2:]

    result = int(s1 + s2) if n != 1 else 0
    return result