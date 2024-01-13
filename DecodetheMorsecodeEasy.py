from preloaded import MORSE_CODE

def decode_morse(morse_code):
    return ' '.join(''.join(MORSE_CODE[k] for k in char.split(' ')) for char in morse_code.strip().split("   "))
    pass