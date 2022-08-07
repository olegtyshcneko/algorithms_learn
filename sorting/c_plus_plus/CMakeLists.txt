cmake_minimum_required(VERSION 3.16)
project(c_plus_plus)

set(CMAKE_CXX_STANDARD 14)

include(FetchContent)
FetchContent_Declare(
        googletest
        # Specify the commit you depend on and update it regularly.
        URL https://github.com/google/googletest/archive/5b909beeec178f338be997830b6c31a80cda7a93.zip
)
FetchContent_MakeAvailable(googletest)

add_executable(c_plus_plus sorting_test.cpp insertion_sort.cpp insertion_sort.h)

target_link_libraries(c_plus_plus GTest::gtest_main)

include(GoogleTest)
gtest_discover_tests(c_plus_plus)