<?php 
function asal_sayi_mi($sayi){
    if ($sayi < 2){ return false;}

    $max_kontrol_sayi = (int) $sayi / 2;

    for($index = 2; $index < $max_kontrol_sayi; $index++){
        if($sayi % $index == 0){
            return false;
        }
    }

    return true;
}

for ($index = 1; $index < 1000000; $index++){
    $sonuc = asal_sayi_mi($index);
    if($sonuc){
        echo $index."\n";
    }
    
}