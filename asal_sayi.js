function asal_sayi_mi(sayi){
    if ( sayi < 2){ return false; }
    const max_kontrol_sayi = Math.floor(sayi / 2)
    for(let index = 2 ; index <= max_kontrol_sayi; index++ ){
        if (sayi % index == 0){
            return false;
        }
    }

    return true
}

for(let index = 2; index < 1000000; index++){
    sonuc = asal_sayi_mi(index)
    if(sonuc){
        console.log(index)
    }
    
}
