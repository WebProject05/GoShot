package split

func ChunkLines(lines []string, maxLines int) [][]string {
	var chunk [][]string

	if maxLines <= 0 {
		return [][]string{lines}
	}

	for i := 0; i < len(lines); i += maxLines {
		end := i + maxLines
		if end > len(lines) {
			end = len(lines)
		}

		chunk = append(chunk, lines[i:end])

	}

	return chunk
}