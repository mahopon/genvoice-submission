import { Button, message } from 'antd'
import React, { useState, useRef, useEffect } from 'react'
import { base64ToBlob } from '../utils/encoding';

interface VoiceRecorderProps {
  audioFile?: string; // Optional prop for an audio file as a Base64-encoded string
}

const VoiceRecorder: React.FC<VoiceRecorderProps> = ({ audioFile }) => {
  const [recording, setRecording] = useState<boolean>(false)
  const [audioUrl, setAudioUrl] = useState<string | null>(null);
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef<Blob[]>([]);

  // If audioFile prop is passed, decode and display the audio file
  useEffect(() => {
    if (audioFile) {
      try {
        const decodedBlob = base64ToBlob(audioFile, 'audio/webm');
        const audioUrl = URL.createObjectURL(decodedBlob);
        setAudioUrl(audioUrl);
      } catch (error) {
        console.error("Error decoding audioFile:", error);
        message.error("Failed to decode audio file");
      }
    }
  }, [audioFile]);

  const handleRecordClick = async () => {
    if (!recording) {
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
          const audioUrl = URL.createObjectURL(audioBlob);
          setAudioUrl(audioUrl);
          message.success('Recording complete');
        };

        mediaRecorder.start();
        setRecording(true);
        message.success('Recording started');
      } catch (err) {
        console.log(err);
        message.error('Microphone access denied or unavailable');
      }
    } else {
      mediaRecorderRef.current?.stop();
      setRecording(false);
    }
  };

  const handleDeleteClick = () => {
    setAudioUrl(null);
  };

  return (
    <>
      <Button
        type={recording ? 'default' : 'primary'}
        danger={recording}
        onClick={handleRecordClick}
      >
        {recording ? "Recording..." : "Record"}
      </Button>
      {(audioUrl || audioFile) && (
        <div style={{ marginTop: '2rem' }}>
          <audio controls src={audioUrl || (audioFile ? URL.createObjectURL(base64ToBlob(audioFile, 'audio/webm')) : '')} />
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
}

export default VoiceRecorder;
