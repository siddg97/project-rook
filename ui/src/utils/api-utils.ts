import axios from "axios";

const client = axios.create();

export function uploadResume(file: File, uid: string) {
    const data = new FormData();
    data.append("file", file);

    client.put(`http://localhost:3000/v1/${uid}/resume`, data, { // TODO: Replace base URL with env
        headers: {
            'Content-Type': 'multipart/form-data',
        }
    });
}
