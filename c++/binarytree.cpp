struct treeNode{
    item_type item;
    treeNode *left,*right,*parent;
};

void insertbinarytree(treeNode* &l, item_type x, treeNode *parent){
    treeNode *tmp; // new treeNode
    
    if(l==NULL){
        tmp=malloc(sizeof(treeNode))
        tmp->item=x;
        tmp->left=tmp->right=NULL;
        tmp->parent=parent;
        l=tmp;
        return;
    }
    
    if(x<(l->item){
        insertbinarytree(l->left, x, l);
    }else{
        insertbinarytree(l->right, x, l);
    }
 }