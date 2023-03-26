import os
import json
import librosa
import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.svm import SVC
from sklearn.naive_bayes import GaussianNB
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

# 1. Clone the GitHub repository
os.system("git clone https://github.com/cricketclub/gridspace-stanford-harper-valley.git")
#!git clone https://github.com/cricketclub/gridspace-stanford-harper-valley.git

# 2. Loop through .wav files and their corresponding .json files
data_path = "gridspace-stanford-harper-valley/data/caller/"
wav_files = [f for f in os.listdir(data_path) if f.endswith(".wav")]
json_files = [f.replace(".wav", ".json") for f in wav_files]

# Function to extract acoustic features
def extract_acoustic_features(audio_file):
    y, sr = librosa.load(audio_file)
    zcr = np.mean(librosa.feature.zero_crossing_rate(y))
    energy = np.mean(librosa.feature.rmse(y))
    spectral_centroid = np.mean(librosa.feature.spectral_centroid(y, sr=sr))
    spectral_spread = np.mean(librosa.feature.spectral_bandwidth(y, sr=sr))
    pitch = np.mean(librosa.feature.tonnetz(y, sr=sr))
    spectral_entropy = np.mean(librosa.feature.spectral_contrast(y, sr=sr))
    return zcr, energy, spectral_centroid, spectral_spread, pitch, spectral_entropy

# 3. Transcribe audio using Whisper (replace with actual Whisper code)
def transcribe_audio_whisper(audio_file):
    # Implement actual Whisper transcription here
    return "transcript"

# 4. Extract linguistic features using Word2Vec, GloVe, and BERT (replace with actual feature extraction code)
def extract_linguistic_features(transcript):
    # Implement actual linguistic feature extraction here
    return "word2vec_features", "glove_features", "bert_features"

# 5. Extract acoustic features and create data
data = []
for wav_file, json_file in zip(wav_files, json_files):
    with open(data_path + json_file) as f:
        transcript = json.load(f)["transcript"]
    
    whisper_transcript = transcribe_audio_whisper(data_path + wav_file)
    word2vec_features, glove_features, bert_features = extract_linguistic_features(whisper_transcript)
    acoustic_features = extract_acoustic_features(data_path + wav_file)
    
    row = (wav_file, transcript, whisper_transcript, word2vec_features, glove_features, bert_features, *acoustic_features)
    data.append(row)

# 6. Export data to a CSV file
column_names = ["Audio", "Transcript", "WhisperTranscript", "Word2Vec", "GloVe", "BERT", "ZCR", "Energy", "SpectralCentroid", "SpectralSpread", "Pitch", "SpectralEntropy"]
df = pd.DataFrame(data, columns=column_names)
df.to_csv("output.csv")