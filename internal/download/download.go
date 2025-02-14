package download

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

func Download(url string, ctx context.Context, w io.Writer) error {
	ytDlpCmd := exec.CommandContext(ctx, "yt-dlp", "-o", "-", url)
	ffmpegCmd := exec.CommandContext(ctx, "ffmpeg", "-i", "-", "-f", "mp3", "-")

	ffmpegCmd.Stdin, _ = ytDlpCmd.StdoutPipe()
	ffmpegStdout, _ := ffmpegCmd.StdoutPipe()

	if err := ytDlpCmd.Start(); err != nil {
		return fmt.Errorf("failed to start ytdlp: %w", err)
	}

	if err := ffmpegCmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	if _, err := io.Copy(w, ffmpegStdout); err != nil {
		return fmt.Errorf("failed to copy to writer: %w", err)
	}

	if err := ytDlpCmd.Wait(); err != nil {
		return fmt.Errorf("yt-dlp failed: %w", err)
	}

	if err := ffmpegCmd.Wait(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w", err)
	}

	return nil
}
