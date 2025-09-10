package main

import "errors"

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("不正な式")
	case number:
		// t.valがnumber型だとわかったので
		// int型の値を返す
		return int(val), nil
	case operator:
		// t.valがoperator型だと分かったので
		// 左右のこの値を取得し
		// operatorのメソッドprocess()を呼び出し
		// 値を処理した結果を返す
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// treeValに新しい型が定義されたが、
		// walkTreeは更新されていないことがわかる
		return 0, errors.New("不明な型のノード")
	}
}
