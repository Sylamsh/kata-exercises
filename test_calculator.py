from calculator import add
import pytest

def test_add_returnInt():
    assert type(add('')) == int

def test_add_emptyInput():
    assert add('') == 0

def test_add_oneInput():
    assert add('1') == 1

def test_add_twoInput():
    assert add('1,5') == 6

def test_add_moreInputs():
    assert add('1,5,2,1,6') == 15

def test_add_separatedNewLine():
    assert add('3\n2') == 5

def test_add_separatedNewLineAndComma():
    assert add('3\n2,1') == 6

def test_add_invalidInput():
    with pytest.raises(Exception) as exc:
        add('3\n2,1,')
    assert "trailing delimeter ',' at 5" in str(exc.value)

def test_add_withSetDelimeter():
    assert add("//:\n3:2:2:57") == 64

def test_add_withSetStringDelimeter():
    assert add("//sep\n3sep2sep2sep57") == 64

def test_add_withSetDelimeterInvalidInput():
    with pytest.raises(Exception) as exc:
        add("//|\n1|2|3,4")
    assert "'|' expected but ',' found at position 5" in str(exc.value)

def test_add_oneNegativeNumber():
    with pytest.raises(Exception) as exc:
        add("1,-2")
    assert "negative number(s) not allowed: -2" in str(exc.value)

def test_add_multipleNegativeNumbers():
    with pytest.raises(Exception) as exc:
        add("1,-2, -3")
    assert "negative number(s) not allowed: -2, -3" in str(exc.value)

def test_add_multipleErrorsRaised():
    with pytest.raises(Exception) as exc:
        add("//|\n1|2,-3")
    assert "negative number(s) not allowed: -3" in str(exc.value)
    assert "'|' expected but ',' found at position 3" in str(exc.value)

def test_add_biggerThanThousand():
    assert add('1001,2') == 2

if __name__ == '__main__':
    retcode = pytest.main()
