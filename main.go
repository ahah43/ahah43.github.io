package main

import(
   "fmt"
    "syscall/js"
    "strconv"
    "math"
    //"sort"
	"gonum.org/v1/gonum/stat/combin"
)
func sum(array []int) int {  
 result := 0  
 for _, v := range array {  
  result += v  
 }  
 return result  
}  
func sumF(array []float64) float64 {  
 result := 0.0 
 for _, v := range array {  
  result += v  
 }  
 return result  
}  
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func t_func(d,th,D float64) float64 {
    return th*math.Ceil(math.Sqrt(D*D/4.0-d*d/4.0)/th)
}


func newParallelMethod(Comb []int,availableSteps []int, D float64, th float64) (float64, []float64, []int){
        //sort.Sort(sort.Reverse(sort.IntSlice(Comb)))
        //fmt.Println(Comb)
        area := 0.0
        ti := 0.0
        t_1 := 0.0
        ti_1 := 0.0
        var Ts []float64
        for _, v := range Comb {
            thisStep:= float64(availableSteps[v])
            //fmt.Println(thisStep)
            ti = t_func(thisStep,th,D)
            ti_1 = ti-t_1
            t_1 = ti
            Ts = append(Ts, ti_1)
            //fmt.Println(ti_1)
            area += 2*ti_1*thisStep
            //fmt.Println(area)
            //fmt.Println("******************")
        }
        return area, Ts, Comb
        }




// add function takes the ids of DOM elements, readstheor values, and lastly operates
func see(this js.Value, i []js.Value) interface{}{
    bitSize := 64 // float precision
    jsDoc := js.Global().Get("document")
    D_js := jsDoc.Call("getElementById", "D").Get("value").String()

UL_js := jsDoc.Call("getElementById", "UL").Get("value").String()
LL_js := jsDoc.Call("getElementById", "LL").Get("value").String()


    th_js := jsDoc.Call("getElementById", "th").Get("value").String()
    SF_js := jsDoc.Call("getElementById", "SF").Get("value").String()
    k_js := jsDoc.Call("getElementById", "k").Get("value").String()


    D,_ := strconv.Atoi(D_js)
UL,_ := strconv.Atoi(UL_js)
LL,_ := strconv.Atoi(LL_js)
    th,_ := strconv.ParseFloat(th_js,bitSize)
    SF,_ := strconv.ParseFloat(SF_js,bitSize)
    k,_ := strconv.Atoi(k_js)
    println("D = ", D)
    println("th = ", th)
    println("SF = ", SF)
    

    
    Steps_js := jsDoc.Call("getElementsByName", "Steps")
    StepsCount := Steps_js.Length() // an integer
    println("Number of all available steps = ", StepsCount)


    var available_steps []int
    for i := 0; i < StepsCount; i++ {
          thisStep_js := Steps_js.Call("item", i)
          thisStep_js.Call("setAttribute","data-t","0.0")
          thisStep_status := thisStep_js.Get("checked").Bool()
            //println(thisStep_status)
          if (thisStep_status){
                thisStepValueString := thisStep_js.Get("value").String() 
                thisStepValue,_ := strconv.Atoi(thisStepValueString)
                //println("this step = ", thisStepValue)
                if (thisStepValue < D && thisStepValue>=LL && thisStepValue<=UL){
                    println(thisStepValue)
                    available_steps = append(available_steps,thisStepValue)
                }
             }
        }
n := len(available_steps)
k = Min(k,n)
println("n = ", n)
println("k = ", k)
//  println("a step = ", available_steps)   
//===========================================================
    gen := combin.NewCombinationGenerator(n, k)
	idx := 1
    bestArea := 0.0
    var bestArrangement[] float64
    var bestCombination[] int
	for gen.Next() {
        thisArea, thisArrangement, thisComb := newParallelMethod(gen.Combination(nil),available_steps, float64(D), th)
        //if ! math.IsNaN(thisArea){
        if (thisArea > bestArea){
            bestArea = thisArea
            bestArrangement = thisArrangement
            bestCombination = thisComb
            }
        //}
		idx++
	}
println("all solutions Count = ", idx)
println("Best Arrangement = ", bestArrangement)
println("Best Combination = ", bestCombination)
//===========================================================
//    jsDoc.Call("getElementById", "result").Set("value", bestArea)
bestAreaStr := fmt.Sprintf("Best Area = <br /> %f mm2",bestArea)
jsDoc.Call("getElementById", "result").Set("innerHTML", bestAreaStr)
thisId := ""
thisT := ""
  for i := 0; i < len(available_steps); i++ {
    thisId = strconv.Itoa(available_steps[i])
    thisT = "0.0"
    jsDoc.Call("getElementById",thisId).Call("setAttribute","data-t",thisT)
    jsDoc.Call("getElementById",thisId).Call("setAttribute","innerHTML",thisId)
    
    }


  for i := 0; i < len(bestCombination); i++ {
    thisId = strconv.Itoa(available_steps[bestCombination[i]])
    thisT = fmt.Sprintf("%f",bestArrangement[i])
    jsDoc.Call("getElementById",thisId).Call("setAttribute","data-t",thisT)
    }
//jsDoc.Call("drawCore")
    return nil
}



func registerCallbacks() {
    js.Global().Set("see", js.FuncOf(see))
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}

