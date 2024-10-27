mod accounts_merge;
mod find_redundant_connection;

struct Solution;

struct UnionFind {
    parent: Vec<usize>,
    cnt: usize,
}

impl UnionFind {
    fn new(n: usize) -> Self {
        let parent = (0..n).collect();
        UnionFind {
            parent: parent,
            cnt: n,
        }
    }

    fn find(&mut self, i: usize) -> usize {
        if i != self.parent[i] {
            self.parent[i] = self.find(self.parent[i]);
        }
        self.parent[i]
    }

    fn union(&mut self, i: usize, j: usize) -> bool {
        let x = self.find(i);
        let y = self.find(j);
        if x == y {
            true
        } else {
            self.parent[x] = y;
            false
        }
    }
}
