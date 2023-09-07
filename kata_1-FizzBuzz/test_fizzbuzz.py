from fizzbuzz import fizzbuzz

def test_fizzbuzz_returnstrnumber():
    assert fizzbuzz(1) == '1'

def test_fizzbuzz_returnFizz():
    assert fizzbuzz(3) == 'Fizz'

def test_fizzbuzz_returnBuzz():
    assert fizzbuzz(25) == 'Buzz'

def test_fizzbuzz_returnFizzBuzz():
    assert fizzbuzz(225) == 'FizzBuzz'