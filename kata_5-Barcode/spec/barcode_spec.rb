require 'barcode'

describe BarCode do
    context '.price' do
        it 'returns $7.25 for the input 12345' do
            expect(BarCode.price('12345')).to eq('$7.25')
        end
        it 'returns $12.50 for the input 23456' do
            expect(BarCode.price('23456')).to eq('$12.50')
        end
        it 'returns error not found for the input 99999' do
            expect { BarCode.price('99999') }.to raise_error("Error: barcode not found")
        end
        it 'returns error empty for the empty input' do
            expect { BarCode.price('') }.to raise_error("Error: empty barcode")
        end
    end
    context '.total' do
        it 'returns $19.75 for the input (12345,23456)' do
            expect(BarCode.total('12345','23456')).to eq('$19.75')
        end
    end
end