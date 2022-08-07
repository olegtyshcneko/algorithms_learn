fn main() {
    println!("Hello, world!");
}

fn sort_insertion(input: &[i32]) -> Vec<i32> {
    let mut vec = Vec::with_capacity(input.len());

    for v in input.iter() {
        if vec.is_empty() {
            vec.push(*v);
            continue;
        }

        for (idx, val) in vec.iter().enumerate() {
            if v <= val {
                vec.insert(idx, *v);
                break;
            }

            if idx == vec.len()-1 {
                vec.push(*v);
                break;
            }
        }
    }

    vec
}

#[cfg(test)]
mod tests {
    use crate::sort_insertion;

    static UNSORTED_ARR: [i32; 11] = [5, 2, 2, 1, 4, 0, 3, -5, 10, 24, 7];
    static SORTED_ARR: [i32; 11] = [-5, 0, 1, 2, 2, 3, 4, 5, 7, 10, 24];

    #[test]
    fn test_insertion_sort() {
        println!("hello tests");
        assert_sort(sort_insertion, &UNSORTED_ARR.to_vec(), &SORTED_ARR.to_vec());
    }

    fn assert_sort(f_sort: fn(&[i32]) -> Vec<i32>, input_vec: &Vec<i32>, expected_vec: &Vec<i32>) {
        let sorted_vec = f_sort(input_vec);
        assert_eq!(&sorted_vec, expected_vec)
    }
}