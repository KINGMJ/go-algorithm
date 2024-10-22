package floodfill

// 深度优先搜索
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	originalColor := image[sr][sc]
	dsf(image, sr, sc, originalColor, color)
	return image
}

func dsf(image [][]int, sr int, sc int, originalColor int, color int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != originalColor {
		return
	}
	image[sr][sc] = color
	dsf(image, sr+1, sc, originalColor, color)
	dsf(image, sr-1, sc, originalColor, color)
	dsf(image, sr, sc+1, originalColor, color)
	dsf(image, sr, sc-1, originalColor, color)
}

// 广度优先搜索，使用队列
func floodFillBFS(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	originalColor := image[sr][sc]
	// 定义一个队列
	queue := [][]int{{sr, sc}}
	// 定义一个方向数组
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		// 出队列
		pixel := queue[0]
		queue = queue[1:]
		r, c := pixel[0], pixel[1]
		if image[r][c] == originalColor {
			image[r][c] = color
			// 将相邻的节点入队列
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if isValid(image, newR, newC) && image[newR][newC] == originalColor {
					queue = append(queue, []int{newR, newC})
				}
			}
		}
	}
	return image
}

func isValid(image [][]int, r, c int) bool {
	return r >= 0 && r < len(image) && c >= 0 && c < len(image[0])
}
