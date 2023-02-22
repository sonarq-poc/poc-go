package main

import (
	"testing"
)

func TestAddition3(t *testing.T) {
    sum := addition3(1,2);
    if sum == 3 {
        t.Log("Recieved as expected: 3")
    } else {
        t.Error(sum)
    }
}

// func TestAddInt(t *testing.T) { 
//     testCases := []struct { 
//         Name     string 
//         Values   []int 
//         Expected int 
//     }{ 
//         {"addInt() -> 0", []int{}, 0}, 
//         {"addInt([]int{10, 20, 100}) -> 130", []int{10, 20, 100}, 130}, 
//     } 
 
//     for _, tc := range testCases { 
//         t.Run(tc.Name, func(t *testing.T) { 
//             sum := addInt(tc.Values...) 
//             if sum != tc.Expected { 
//                 t.Errorf("%d != %d", sum, tc.Expected) 
//             } else { 
//                 t.Logf("%d == %d", sum, tc.Expected) 
//             } 
//         }) 
//     } 
// }