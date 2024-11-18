mod cal_points;
mod check_record;
mod difference_of_sum;
mod find_winning_player;
mod image_smoother;
mod judge_square_sum;
mod losing_player;
mod max_height_of_triangle;
mod max_score_sightseeing_pair;
mod min_operations;
mod minimum_average;
mod neigh_bor_sum;
mod number_of_pairs;
mod results_array;
mod time_required_to_buy;

struct Solution;

#[derive(PartialEq, Eq, Clone, Debug)]
struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
