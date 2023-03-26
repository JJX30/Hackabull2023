import os
import csv
import whisper
from datasets import load_dataset

model = whisper.load_model("base")

# Define the path to the dataset on the USB drive
dataset_path = '/Volumes/USB322FD/data'

# Specify the output CSV file on the USB drive
output_file = 'output.csv'
output_path = os.path.join(dataset_path, output_file)

# Load the dataset from the specified data directory
dataset = load_dataset('MLCommons/peoples_speech', split='train', cache_dir=dataset_path)

for example in dataset['train']:
    if example['audio_file'].endswith('.flac'):
        audio = whisper.load_audio(example['audio_file'])
        audio = whisper.pad_or_trim(audio)

        # Transcribe speech
        mel = whisper.log_mel_spectrogram(audio).to(model.device)
        result = whisper.decode(model, mel)

        transcription = result[0].text

        # Write row to CSV file on the USB drive
        with open(output_path, 'a', newline='') as csvfile:
            csvwriter = csv.writer(csvfile)
            csvwriter.writerow([example['audio_file'], transcription])

