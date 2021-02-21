中序遍历的递归版本是

```shell script
dfs(root.left)
打印节点 root
dfs(root.right)
```


//前序遍历
void preorder(TreeNode *root, vector<int> &path)
{
    if(root != NULL)
    {
        path.push_back(root->val);
        preorder(root->left, path);
        preorder(root->right, path);
    }
}


//中序遍历
void inorder(TreeNode *root, vector<int> &path)
{
    if(root != NULL)
    {
        inorder(root->left, path);
        path.push_back(root->val);
        inorder(root->right, path);
    }
}


//后续遍历
void postorder(TreeNode *root, vector<int> &path)
{
    if(root != NULL)
    {
        postorder(root->left, path);
        postorder(root->right, path);
        path.push_back(root->val);
    }
