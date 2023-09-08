class BarCode
    def self.price_float(code)
        if code == '12345'
            7.25
        elsif code == '23456'
            12.50   
        elsif code == '99999'
            raise "Error: barcode not found"
        elsif code == ''
            raise "Error: empty barcode"
        end
    end
    def self.price(code)
        "$%0.2f" % [price_float(code)]
    end
    def self.total(*codes)
        total = 0.0
        codes.each do |code|
            total = total + price_float(code)
        end
        "$%0.2f" % [total]
    end
end