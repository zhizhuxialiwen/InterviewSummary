# 树

## 1、红黑树

红黑树

蒙德里安的梦想 2020-02-22 15:05:47  1629  收藏 2
分类专栏： C++
版权

### 1.1、目录：

红黑树概念
红黑树性质
红黑树的定义
红黑树的插入(重点）
情况1（违反原因、解决方法、整体图解）
情况2（情况的产生、解决方法、整体图解、方法判断）
情况3（方法判断、解决方法、整体图解）
三种情况的总结
红黑树的代码实现
讲前小结：红黑树的插入部分理解的时候要多画图，搞清楚不同情况下的插入逻辑。其实画图花多了，理解起来就不是那么的费劲。
本节大部分是图解。红黑树中的红色节点用红色代替，黑色节点用黑色代替

### 1.2、红黑树的概念

红黑树，其实就是一种二叉搜索树的优化。每一个节点都会有一个颜色进行标记（红色，或者黑色）。它的五条性质决定它是一种近似平衡的二叉树。

### 1.3、红黑树的性质

1. 每各节点不是黑色就是红色
2. 根节点是黑色
3. 如果一个节点是红色的，则它的两个孩子节点是黑色的
4. 对于每个节点，从节点到其所有后代叶节点的路径上，包含相同数目的黑色节点
5. 每个叶子节点都是黑色的（叶子节点是指空节点）

我们根据上面五条性质，可以知道红黑树中最主要的性质是：

红黑树中不会出现连续的红节点（第3条性质）
每一条路径上的黑色节点的数量相同（第4条性质）
还能由这五条性质可以推出：
红黑树中最长路径中节点个数不会超过最短路径节点个数的两倍
假设每一条路径上有X个黑色节点，红黑树中红色节点，不能连续出现，所以可以穿插在黑色节点，肯定不会破坏性质4. 那么最长的路径就是2x-1. 2x - 1 < 2*x 。所以最长路径中节点的个数不会超过最短路径中节点个数的两倍。


### 1.4、红黑树节点的定义

红黑树是三叉链型的结构。所以每一个节点中都要有三个指针，指向左右子节点和父亲节点

```c++
enum Colour //标记红黑树的节点的颜色
{
	BLACK,
	RED,
};

// 红黑的节点
template <class T>
struct RBTreeNode
{
	RBTreeNode<T>* _left; 
	RBTreeNode<T>* _right;
	RBTreeNode<T>* _parent;

	T _data;   // 节点的值
	Colour _col; // 节点的颜色

	RBTreeNode(const T& data)  //节点的构造函数
		:_left(nullptr)
		, _right(nullptr)
		, _parent(nullptr) 
		, _data(data)
		, _col(RED)
	{}
};
```

红黑树中最重要的步骤就是红黑树的插入

### 1.5、红黑树的插入

红黑树是对二叉搜索数的优化，所以插入节点的要求还是按照二叉搜索树的，小于节点的，插入到左子树，大于节点的，插入到右子树
红黑树中由于有五条性质限制，所以我们要考虑插入节点是否破坏了当前的红黑树，而且对于破坏结构的红黑树，还要用 调整红黑树中节点的颜色 和旋转的方法适当调整红黑树，使红黑树满足五条性质。

- 我们在插入的时候插入节点选择红色还是黑色?

我们插入的新节点也应该尽量的不破坏红黑树的五条性质。所以我们插入的新节点默认颜色是红色。如果我们插入的节点是黑色，则一定会破坏性质4，使得每一条路径上黑色节点的数目不一样多。 如果我们插入的节点是红色，我们有可能破坏性质3，如果父节点也是红色的，那么我们就对红黑树进行调节。所以插入的时候，默认新节点的颜色是红色

但是插入之后，就有可能违反红黑树的五条性质，破坏红黑树的结构。所以我们要分情况来讨论。

总的特殊情况分为三类
约定cur为当前节点，p为父节点(parent)，g为祖父节点(grandparent)，u为叔叔节点(uncle)

- 情况一： cur为红色，parent为红色，grandparent为黑色，uncle存在且为红色

违反原因：
插入的cur为红色节点，parent节点也是红色，两个红色节点连在一起，违反了第三条性质，破坏了原本的红黑树结构。

解决方法：
将parent节点和uncle节点改成黑色，grandfather节点改为红色。然后把grandfather当成cur,继续向上调整
如下图（cur插入在parent的左右 或者 parent和uncle互换左右都是一样的）

整体图解：

![tree1](../images/tree1.PNG)

在这里插入图片描述

注意 ：grandparent节点不是根节点，但是当grandparent的上一个节点还是红节点的话，还得继续向上调整。

![tree2](../images/tree2.PNG)
在这里插入图片描述

- 情况二：cur为红色，parent为红色，grandparent为黑色，uncle不存在/uncle为黑色

情况二的产生
根据uncle节点的有无，cur节点的又有两种不同的解释（解决的方案都一样，而且弄懂逻辑就很简单了）

当uncle节点不存在时，则cur一定是新插入的节点，因为每一条路径上的黑色节点个数要相同。

![tree3](../images/tree3.PNG)

在这里插入图片描述

当uncle节点存在且黑为色是，cur节点的原来的颜色一定是为黑色的。红色的原因是cur的子树在调整的过程正将cur节点的颜色由黑色改为红色。

![tree4](../images/tree4.PNG)

在这里插入图片描述
解决方法：
parent为grandparen的左孩子，cur为p的左孩子，则进行右单旋
parent为grandparent的右孩子，cur为p的右孩子，则进行左单旋
parent变成黑色
grandparent变为红色。

整体图解：

![tree5](../images/tree5.PNG)
在这里插入图片描述

情况二判断旋转的方法：
在情况二中，我们可以看到

cur、parent、grandparent都是在一条线上的。
左边高，将左边压下来，所以进行右单旋。
右边高，将右边压下来，所以进行左单旋
例如：

![tree6](../images/tree6.PNG)

在这里插入图片描述

- 情况三：cur为红色，parent为红色，grandparent为黑色，uncle不存在/uncle为黑色
方法判断
情况三就是情况二的特例，只不过是cur、parent、grandparent三者不在同一条线上
我么就是通过将情况三转换到情况二做。

![tree7](../images/tree7.PNG)

在这里插入图片描述

解决方法：

当parent为grandparent的左孩子时，cur为parent的右孩子，则对parent进行左单旋 (上图第一种形态）。
当parent为grandparent的右孩子时，cur为parent的左孩子，则对parent做右单旋（上图第二种形态）。
就转换到了情况二。 针对情况二的解法进行旋转，调节点颜色。
但是情况2还要进行一次旋转，所以就是下面总结的，直接调整完毕。 不是左右双旋 就是 右左双旋

其实就可以总结为： 情况三，直接进行两次旋转：

【左右双旋】当parent为grandparent的左孩子时，cur为parent的右孩子，对parent进行左单旋，得到情况二，再对grandfather进行右单旋
【右左双旋】当parent为grandparent的右孩子时，cur为parent的左孩子，则对parent做右单旋，得到情况二，再对grandfather进行左单旋

### 1.6、三种插入情况的总结

情况一：插入之后，将祖父节点改为红色之后，祖父节点不一定是根节点，有可能还得影响祖父的父亲节点，所以要循环处理
情况二：插入节点，调节颜色后，祖父节点是黑色，不会影响上面节点。而且每一条路径上黑色节点的数目是相同的，不需要向上循环处理
情况三：插入节点，进行左右双旋或者右左双旋之后，祖父节点是黑色，不会影响上面的节点，而且每一条路径上黑色节点相同，不需要向上循环处理

### 1.7、代码实现部分

有关右旋、左旋、和双旋的理解 在VAL树的旋转有讲解

```c++
#pragma once
#include <iostream>
using namespace std;

enum Colour
{
	BLACK,
	RED,
};

template <class T>
struct RBTreeNode
{
	RBTreeNode<T>* _left;
	RBTreeNode<T>* _right;
	RBTreeNode<T>* _parent;

	T _data;
	Colour _col;

	RBTreeNode(const T& data)
		:_left(nullptr)
		, _right(nullptr)
		, _parent(nullptr)
		, _data(data)
		, _col(RED)
	{}
};


template <class K, class V>
class RBTree
{
	typedef RBTreeNode<pair<K, V>> Node;
public:
	RBTree()
		:_root(nullptr)
	{}
	pair<Node*, bool> Insert(const pair<K, V>& kv)
	{
		if (_root == nullptr)
		{
			_root = new Node(kv);
			_root->_col = BLACK;
			return make_pair(_root, true);
		}

		Node* parent = nullptr;
		Node* cur = _root;
		while (cur)
		{
			if (cur->_data.first < kv.first)
			{
				parent = cur;
				cur = cur->_right;
			}
			else if (cur->_data.first > kv.first)
			{
				parent = cur;
				cur = cur->_left;
			}
			else
			{
				return make_pair(cur, false);
			}
		}

		Node* newnode = new Node(kv);
		cur = newnode;
		cur->_col = RED;

		
		if (parent->_data.first < kv.first)
		{
			parent->_right = cur;
			cur->_parent = parent;
		}
		else
		{
			parent->_left = cur;
			cur->_parent = parent;
		}

		//调整红黑树节点颜色
		while (parent && parent->_col == RED) //父亲节点存在 且父亲节点为红 开始调整节点颜色
		{
			//关键看叔叔
			Node* grandfather = parent->_parent;
			if (parent == grandfather->_left)         //如果插入节点cur的父亲节点是祖父的左节点， 那么叔叔节点肯定是祖父的右节点
			{
				Node* uncle = grandfather->_right;
				//情况一
				if (uncle && uncle->_col == RED)
				{
					parent->_col = uncle->_col = BLACK;
					grandfather->_col = RED;


					//向上调整
					cur = grandfather;
					parent = cur->_parent;
				}

				

				else //情况二 + 情况三  uncle不存在或者存在且为黑
				{
				//情况二
					//    g
					//  p
					//c
					if (cur == parent->_left) //grandfather 节点肯定为黑，因为parent节点一直为红
					{
					//右旋
						RotateR(grandfather);
						grandfather->_col = RED;
						parent->_col = BLACK;

					}
				// 情况三 左右双旋
					//    g
					//  p
					//    c
					else  //cur == parent->_right  双旋
					{
						RotateLR(grandfather);
						cur->_col = BLACK;
						grandfather->_col = RED;
					}
					//跳出循环
					break;

				}
			}
			else//如果插入节点cur在父节点的右边的
			{
				Node* uncle = grandfather->_left;  //叔叔节点是父节点的右边
				//情况一  uncle存在且为红
				if (uncle && uncle->_col == RED)
				{
					uncle->_col = parent->_col = BLACK;
					grandfather->_col = RED;

					cur = grandfather;
					parent = cur->_parent;
				}
				else 
				{	//旋转 + 变色   
					//uncle不存在 或 存在且为黑
					//情况二
					// g
					//    p
					//       cur
					if (cur == parent->_right) //单旋 左低右高 左单旋
					{
						RotateL(grandfather);
						parent->_col = BLACK;
						grandfather->_col = RED;
					}
					//情况三
					else  //cur插入在 parent->_left
					{
						// g
						//   p
						// c
						//右左双旋
						RotateRL(grandfather);
						cur->_col = BLACK;
						grandfather->_col = RED;
					}

					break;
				}
				
			}
		}
		_root->_col = BLACK;

		return make_pair(cur, true);
	}
	//左旋
	void RotateL(Node* parent)
	{
		Node* subR = parent->_right;
		Node* subRL = subR->_left;

		parent->_right = subRL;
		if (subRL)
			subRL->_parent = parent;

		Node* Pparent = parent->_parent;
		parent->_parent = subR;
		subR->_left = parent;
		

		if (_root == parent)
		{
			_root = subR;
			subR->_parent = NULL;
		}
		else
		{
			if (Pparent)
			{
				if (Pparent->_left == parent)
				{
					Pparent->_left = subR;
				}
				else
				{
					Pparent->_right = suR;
				}
			}
			subR->_parent = Pparent;
		}
		
		
	}

	void RotateR(Node* parent)
	{
		Node* subL = parent->_left;
		Node* subLR = subL->_right;

		parent->_left = subLR;
		if (subLR)
		{
			subLR->_parent = parent;
		}

		Node* Pparent = parent->_parent;
		subL->_right = parent;
		parent->_parent = subL;
	

		if (parent == _root) //根节点的情况
		{
			_root = subL;
			subL->_parent = nullptr;
		}
		else //非根节点的情况
		{
			if (Pparent->_left == parent)
			{
				Pparent->_left = subL;
			}
			else
			{
				Pparent->_right = subL;
			}
			subL->_parent = Pparent;
		}
		
	}
	void RotateRL(Node* parent) //右左双旋 
	{
		RotateR(parent->_right);
		RotateL(parent);
	}
	void RotateLR(Node* parent)
	{
		RotateL(parent->_left);
		RotateR(parent);
	}
private:
	Node* _root;
};
```