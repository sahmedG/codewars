def decode_bits(bits):
    new_bits = bits
    if new_bits[0] == '0':
        i = 1
        while new_bits[i-1] == '0':
            i+=1
        new_bits = new_bits.strip('0'*i).rstrip('0'*i)
    if new_bits.find('1') != -1 and new_bits.find('0') != -1:
        bits_0 = sorted([len(j) for j in new_bits.split('1') if j !=''])[0]
        bits_1 = sorted([len(j) for j in new_bits.split('0') if j !=''])[0]
        min = (bits_0 if bits_0<=bits_1 else bits_1)
        if min%9 == 0:
            speed = min/3
        else:
            s = min
    elif new_bits.find('1') == -1:
        return ''
    elif new_bits.find('0') == -1:
        return '.'
    return new_bits.replace('111'*int(s), '-').replace('000'*int(s), ' ').replace('1'*int(s), '.').replace('0'*int(s), '')

def decode_morse(morseCode):
    res = ''
    morseCode = morseCode.split(' ')
    lookup = dict.keys(MORSE_CODE)
    for char in morseCode:
        if char in lookup and char != '':
            res += MORSE_CODE[char]
        elif char == '':
            res += ' '
    return res