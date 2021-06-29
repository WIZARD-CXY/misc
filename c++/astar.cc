#include <iostream>
#include <algorithm>
#include <list>
#include <memory>
#include <vector>
using namespace std;
struct PathNode {
   PathNode(const int row, const int col, 
            const int f = numeric_limits<int>::max(),
            const shared_ptr<PathNode> &parent = nullptr)
           : row_(row), col_(col), f_(f), parent_(parent){}

   bool operator==(const PathNode &other) const {
        return row_ == other.row_ && col_ == other.col_;
   }

   int row_,col_;
   int f_; // f=g+h
   shared_ptr<PathNode> parent_;
};

void AStarSearch(const vector<vector<int>> &maze_mat,
                 list<shared_ptr<PathNode>> &solution);

int main() {
    int row,col;
    while(cin>>row && cin>>col) {
      vector<vector<int>> maze_mat(row, vector<int>(col));
      for(int i = 0; i < row; i++) {
        for(int j=0; j < col; j++) {
           cin>>maze_mat[i][j];
        }
      }


      list<shared_ptr<PathNode>> solution;
      AStarSearch(maze_mat, solution);
      for(const auto &node : solution) {
        cout<<"("<<node->row_<<","<<node->col_<<")"<<endl;
      }
    }
    
    return 0;
}

void AStarSearch(const vector<vector<int>> &maze_mat,
                 list<shared_ptr<PathNode>> &solution) {
  if(maze_mat.empty() || maze_mat[0].empty()) {
      return;
  }

  int start_row=0;
  int start_col=0;
  int end_row=maze_mat.size()-1;
  int end_col=maze_mat[0].size()-1;

  list<shared_ptr<PathNode>> open_list, closed_list;
  shared_ptr<PathNode> end_node=nullptr;
  auto start_node = make_shared<PathNode>(start_row,start_col);

  open_list.emplace_back(start_node);

  auto calc_f_val_func = [&start_row, &start_col, &end_row, &end_col](
                            const int row, const int col) {
      int g_val = abs(row-start_row) + abs(col-start_col);
      int h_val = abs(end_row-row) + abs(end_col - col);

      int h_coeff = 3;
      return g_val + h_coeff * h_val;
  };

  auto not_exist_func = [](const list<shared_ptr<PathNode>> &list,
                            shared_ptr<PathNode> &node) {
      return find(list.begin(), list.end(), node) == list.end();
  };


  auto handdle_child_node_func = [&open_list, &closed_list, &start_row,
                                  &start_col, &end_row, &end_col, &maze_mat,
                                  &calc_f_val_func, &not_exist_func](
                                      const shared_ptr<PathNode> &cur_node,
                                      const int row, const int col) {
    if ( row < start_row || col < start_col || row > end_row || col > end_col) {
      return;
    }

    // wall
    if (maze_mat[row][col]>0){
        return;
    }

    auto child_node = make_shared<PathNode>(row, col);
    if(not_exist_func(open_list,child_node) &&
        not_exist_func(closed_list,child_node)) {
        child_node->f_ = calc_f_val_func(row, col);
        child_node->parent_ = cur_node;
        open_list.emplace_back(child_node);
    }
  };


  while(!open_list.empty()) {
    auto min_iter = min_element(open_list.begin(),open_list.end(), 
    [](const shared_ptr<PathNode> &lhs, shared_ptr<PathNode> &rhs){
      return lhs->f_ < rhs->f_;
    });

    if(min_iter == open_list.end() ){
      break;
    }

    auto min_node = *min_iter;
    cout << "min_node: " << min_node << ", " << min_node->row_ << ", "
              << min_node->col_ << ", " << min_node->f_ << std::endl;

    closed_list.emplace_back(min_node);

    open_list.erase(min_iter);

    if(min_node->row_ == end_row && min_node->col_ == end_col){
      end_node=min_node;
      break;
    }else{
      // handle up down left right four directions
      handdle_child_node_func(min_node, min_node->row_ -1,min_node->col_);
      handdle_child_node_func(min_node, min_node->row_+1,min_node->col_);
      handdle_child_node_func(min_node, min_node->row_,min_node->col_+1);
      handdle_child_node_func(min_node, min_node->row_,min_node->col_-1);
    }
  }

  cout<<end_node<<endl;
  if (end_node != nullptr){
    do {
      solution.emplace_back(end_node);
      end_node = end_node->parent_;
    }while(end_node != nullptr);

    solution.reverse();
  }
}
