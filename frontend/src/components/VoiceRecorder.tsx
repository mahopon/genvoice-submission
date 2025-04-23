import { Button, message } from 'antd';
import React, { useState, useRef, useEffect } from 'react';

interface VoiceRecorderProps {
  questionId: string;
  existingAudio?: Blob;
  isRecording: boolean;  // Check if this recorder is the one currently recording
  isDisabled: boolean;    // Disable all other recorders except the active one
  onRecordingComplete?: (audioBlob: Blob) => void;
  onDelete?: (questionId: string) => void;
  onStartRecording?: () => void;  // Callback to start recording for this question
}

const VoiceRecorder: React.FC<VoiceRecorderProps> = ({
  questionId,
  existingAudio,
  isRecording,
  isDisabled,
  onRecordingComplete,
  onDelete,
  onStartRecording
}) => {
  const [audioUrl, setAudioUrl] = useState<string | null>(null);
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef<Blob[]>([]);

  useEffect(() => {
    if (existingAudio) {
      const audioUrl = URL.createObjectURL(existingAudio);
      setAudioUrl(audioUrl);
    }
  }, [existingAudio]);

  const handleRecordClick = async () => {
    if (!isRecording) {
      try {
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        const mediaRecorder = new MediaRecorder(stream);
        mediaRecorderRef.current = mediaRecorder;
        audioChunksRef.current = [];

        mediaRecorder.ondataavailable = event => {
          audioChunksRef.current.push(event.data);
        };

        mediaRecorder.onstop = () => {
          const audioBlob = new Blob(audioChunksRef.current, { type: 'audio/webm' });
          onRecordingComplete?.(audioBlob);
          const audioUrl = URL.createObjectURL(audioBlob);
          setAudioUrl(audioUrl);
          message.success('Recording complete');
        };

        mediaRecorder.start();
        onStartRecording?.(); // Notify parent that recording started
        message.success('Recording started');
      } catch (err) {
        console.log(err);
        message.error('Microphone access denied or unavailable');
      }
    } else {
      mediaRecorderRef.current?.stop();
      message.success('Recording stopped');
    }
  };

  const handleDeleteClick = () => {
    setAudioUrl(null);
    onDelete!(questionId);
  };

  return (
    <>
      <Button
        type={isRecording ? 'default' : 'primary'}
        danger={isRecording}
        onClick={handleRecordClick}
        disabled={isDisabled}  // Disable this recorder if it's not the one recording
      >
        {isRecording ? "Recording..." : "Record"}
      </Button>
      {(audioUrl) && (
        <div style={{ marginTop: '2rem' }}>
          <audio controls src={audioUrl} />
          <br />
          <a
            onClick={handleDeleteClick}
            style={{ color: 'red', cursor: 'pointer' }}
          >
            Delete Recording
          </a>
        </div>
      )}
    </>
  );
};

export default VoiceRecorder;
