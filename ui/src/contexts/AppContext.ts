import { createContext } from "react";
import { initializeApp, FirebaseApp, FirebaseOptions } from 'firebase/app';

export const AppContext = createContext({});

export function initializeFirebase(): FirebaseApp {
    console.log("Initializing Firebase")

    const firebaseConfig: FirebaseOptions = {
        projectId: import.meta.env.FIREBASE_PROJECT_ID,
        apiKey: import.meta.env.FIREBASE_API_KEY,
    };

    return initializeApp(firebaseConfig);
}
