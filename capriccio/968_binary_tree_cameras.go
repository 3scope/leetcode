package main

type cameraStatus int

const (
	noCoverage  cameraStatus = 0
	hasCamera   cameraStatus = 1
	hasCoverage cameraStatus = 2
)

func minCameraCover(root *TreeNode) int {
	result := 0
	rootStatus := minCameraCoverTraversal(root, &result)
	if rootStatus == noCoverage {
		result++
	}

	return result
}

func minCameraCoverTraversal(root *TreeNode, result *int) cameraStatus {
	if root == nil {
		return hasCoverage
	}

	left := minCameraCoverTraversal(root.Left, result)
	right := minCameraCoverTraversal(root.Right, result)
	// left == 0 && right == 0 Both no coverage.
	// left == 1 && right == 0
	// left == 0 && right == 1
	// left == 0 && right == 2
	// left == 2 && right == 0
	if left == noCoverage || right == noCoverage {
		(*result)++
		return hasCamera
	} else if left == hasCamera || right == hasCamera {
		// left == 1 && right == 1 Both has the camera.
		// left == 1 && right == 2
		// left == 2 && right == 1
		return hasCoverage
	} else {
		// left == hasCoverage && right == hasCoverage
		return noCoverage
	}
}
