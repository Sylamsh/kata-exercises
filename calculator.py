def add(s:str) -> int:
    if is_empty(s):
        return 0
    delimeter, sum_expression = ',', s
    if has_custom_delimeter(s):
        delimeter, sum_expression = get_custom_delimeter(s)
    else:
        sum_expression = replace_newline(s, ',')

    check_for_errors(sum_expression, delimeter)

    numbers, answer = sum_expression.split(delimeter), 0
    for i, num in enumerate(numbers):
        if int(num) < 1000:
            answer += int(num)
    return answer


def is_empty(s:str):
    return s == ''

def get_custom_delimeter(s: str) -> (str, str):
    expressions = s.split('\n')
    delimeter = expressions[0][2:]
    sum_expression = expressions[1]
    return delimeter, sum_expression

def replace_newline(s: str, delimeter: str) -> (str):
    return s.replace('\n', delimeter)

def has_custom_delimeter(s:str) -> bool:
    return s[:2] == '//'

def get_unexpected_symbols(s:str, delimeter:str) -> str:
    return s.replace(delimeter, "").replace('0', "").replace('1', "").replace('2', "").replace('3', "").replace('4', "").replace('5', "").replace('6', "").replace('7', "").replace('8', "").replace('9', "")

def get_all_negative_numbers(unexpected_symbols:str, sum_expression:str) -> str:
    neg_nums, separator = '', ', '
    if '-' in unexpected_symbols:
        for i, char in enumerate(sum_expression):
            if char == '-':
                neg_nums += sum_expression[i:i+2] + separator
    return neg_nums

def check_for_errors(sum_expression:str, delimeter:str):
    # error handling
    exceptions = []
    unexpected_symbols = get_unexpected_symbols(sum_expression, delimeter)
    if unexpected_symbols != '':
        #negative numbers
        neg_nums = get_all_negative_numbers(unexpected_symbols, sum_expression)
        if neg_nums != '':  
            exceptions.append(Exception("negative number(s) not allowed: %s"%(neg_nums[:-2])))
        #unexpected tokens
        for symbol in unexpected_symbols:
            if symbol != '-':
                exceptions.append(Exception("'%s' expected but '%s' found at position %d"%(delimeter, symbol, sum_expression.index(symbol))))
    #trailing delimiter
    if sum_expression[-1] == delimeter:
        exceptions.append(Exception("trailing delimeter '%s' at %d"%(delimeter, len(sum_expression)-1)))
    #raise all exceptions if they exist
    if len(exceptions) > 0:
        raise Exception(exceptions)