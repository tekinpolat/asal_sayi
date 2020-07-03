public class AsalSayi {
    public static void main(String[] args) {
        
        for (int index = 2; index < 1_000_000; index++){
            boolean sonuc  = asal_sayi_mi(index);
            if(sonuc){
                System.out.println(index);
            }
        }
        
    }

    public static boolean asal_sayi_mi(int sayi){
        int max_kontrol_sayi = (int)sayi/2;
        for (int index = 2; index < max_kontrol_sayi; index++){
            if (sayi % index == 0){
                return false;
            }
        }
        return true;
    }
}