//
// Created by lotie on 07.08.2022.
//
#include "insertion_sort.h"

using namespace std;

vector<int> sort_insertion(vector<int> const &input) {
    vector<int> sortedArr(input.begin(),  input.end());

    for(auto i = 0; i<input.size()-1; i++) {
        for(auto j = i+1; j > 0; j--) {
            if(sortedArr[j-1] > sortedArr[j]) {
                swap(sortedArr[j-1], sortedArr[j]);
            }
        }
    }

    return sortedArr;
}