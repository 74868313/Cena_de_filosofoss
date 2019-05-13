package main
import "fmt"
var Filosofos_que_comieron int = 1
type Palillos struct{
	ocupado int
}
var Palilloses[5] Palillos


type filosofo struct{
	Izq int
	Der int
}
var Filosofos[5] filosofo

func Comer(pasa_filosofo int){
	var n int = 5
	if Filosofos[pasa_filosofo].Izq==3 && Filosofos[pasa_filosofo].Der==3{
		fmt.Println("\tFilosofo ", (pasa_filosofo+1)," Acabo de comer\n")
	}
	if Filosofos[pasa_filosofo].Izq==1 && Filosofos[pasa_filosofo].Der==1{


		fmt.Println("\tFilosofo ",(pasa_filosofo+1)," Acabo de comer\n")


		Filosofos[pasa_filosofo].Izq = 3
		Filosofos[pasa_filosofo].Der = 3
		var otroPalillos int = pasa_filosofo-1
		if otroPalillos == (-1) {
			otroPalillos=(n-1);
		}
		Palilloses[pasa_filosofo].ocupado = 0
		Palilloses[otroPalillos].ocupado = 0
		fmt.Println("\tFilosofo ",(pasa_filosofo+1),": Dejo Palillos ",(pasa_filosofo+1)," y Palillos ",(otroPalillos+1),"\n")
		Filosofos_que_comieron++;
	}
	if Filosofos[pasa_filosofo].Izq==1 && Filosofos[pasa_filosofo].Der==0 {
		if pasa_filosofo==(n-1) {
			if Palilloses[pasa_filosofo].ocupado==0 {
				Palilloses[pasa_filosofo].ocupado = 1
				Filosofos[pasa_filosofo].Der = 1
				fmt.Println("\tPalillos ",(pasa_filosofo+1)," ocupado por Filosofo ",(pasa_filosofo+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(pasa_filosofo+1)," esta esperando para Palillos ",(pasa_filosofo+1),"\n")
			}
		}else{
			var IDdublicado int = pasa_filosofo ;
			pasa_filosofo = pasa_filosofo-1;


			if pasa_filosofo== -1{
				pasa_filosofo =(n-1);
			}
			if Palilloses[pasa_filosofo].ocupado == 0{
				Palilloses[pasa_filosofo].ocupado = 1
				Filosofos[IDdublicado].Der = 1
				fmt.Println("\tPalillos ",(pasa_filosofo+1)," ocupado por Filosofo ",(IDdublicado+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(IDdublicado+1)," esta esperando para Palillos ",(pasa_filosofo+1),"\n")
			}
		}
	}
	if Filosofos[pasa_filosofo].Izq==0 {
		if pasa_filosofo==(n-1) {
			if Palilloses[pasa_filosofo-1].ocupado == 0 {
				Palilloses[pasa_filosofo-1].ocupado = 1
				Filosofos[pasa_filosofo].Izq = 1
				fmt.Println("\tPalillos ",pasa_filosofo," ocupado por Filosofo ",(pasa_filosofo+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(pasa_filosofo+1)," esta esperando por Palillos ",pasa_filosofo,"\n")
			}
		}else{
			if Palilloses[pasa_filosofo].ocupado == 0{
				Palilloses[pasa_filosofo].ocupado = 1
				Filosofos[pasa_filosofo].Izq = 1
				fmt.Println("\tPalillos ",(pasa_filosofo+1)," ocupado por Filosofo ",(pasa_filosofo+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(pasa_filosofo+1)," esta esperando para Palillos ",(pasa_filosofo+1),"\n")
			}
		}
	}


}
func main(){
	var n int = 5
	for i:=0; i<n ;i++{
		Palilloses[i].ocupado = 0
		Filosofos[i].Izq = 0
		Filosofos[i].Der = 0
	}
	for Filosofos_que_comieron<n{
		for i:=0;i<n;i++ {
			Comer(i);
		}
		fmt.Println("\t hasta ahora el numero de filosofo que Termino la cena es ",Filosofos_que_comieron,"\n\n")
	}
}

