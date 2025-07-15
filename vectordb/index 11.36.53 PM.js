import dotenv from 'dotenv'
import { GoogleGenAI } from "@google/genai";

dotenv.config();

async function main() {
    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    try {
        const response = await ai.models.embedContent({
            model: 'gemini-embedding-exp-03-07', 
            contents: {
                parts: [
                    { text: 'what is the meaning of life' }
                ]
            },
        });
        return response.embeddings;
    } catch (error) {
        console.error("Error embedding content:", error);
        if (error.response) {
            console.error("API Error Response:", error.response.status, error.response.data);
        }
    }
}

const embeddings = await main();
console.log(embeddings[0].values);  

