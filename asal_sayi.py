def asal_sayi_mi(sayi):
    if sayi < 2: return False
    
    max_kontrol_sayi = sayi//2                      #girilen sayının yarısına kadar kontrol etmek yeterli
    for index in range(2, max_kontrol_sayi + 1):    #2 den başlayıp sayısının yarısına kadar 
        if sayi % index == 0:
            return False

    return True 

for index in range(1, 1_000_000):
    sonuc   = asal_sayi_mi(index)
    if sonuc:
        print("{:,}".format(index))