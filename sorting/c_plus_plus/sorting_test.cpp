//
// Created by lotie on 07.08.2022.
//

#include <gtest/gtest.h>
#include "insertion_sort.h"

const std::vector<int> unsortedArr{5, 2, 2, 1, 4, 0, 3, -5, 10, 24, 7};
const std::vector<int> sortedArr{-5, 0, 1, 2, 2, 3, 4, 5, 7, 10, 24};

TEST(SortingTests, InsertionSortTest) {
    auto newSortedArr = sort_insertion(unsortedArr);

    ASSERT_EQ(newSortedArr.size(), sortedArr.size());

    for (auto i = 0; i < newSortedArr.size(); i++) {
        EXPECT_EQ(newSortedArr[i], sortedArr[i])
            << "Value differ at index "
            << i
            << " Expected val: "
            << sortedArr[i]
            << " Actual val: "
            << newSortedArr[i]
            << std::endl;
    }
}