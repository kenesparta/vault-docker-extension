import React from 'react';
import Button from '@mui/material/Button';
import {createDockerDesktopClient} from '@docker/extension-api-client';
import {Stack, TextField, Typography} from '@mui/material';

// Note: This line relies on Docker Desktop's presence as a host application.
// If you're running this React app in a browser, it won't work properly.
const client = createDockerDesktopClient();

function useDockerDesktopClient() {
    return client;
}

async function fetchData(url: string): Promise<any> {
    try {
        const response = await fetch(url);

        // Check if the response is OK (status code in the range 200-299)
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        // Parse JSON data from response
        return await response.json(); // Handle the data
    } catch (error) {
        // Handle any errors
        console.error('Fetch error:', error);
        throw error; // Rethrow the error for further handling if necessary
    }
}

export function App() {
    const [response, setResponse] = React.useState<string>();
    const ddClient = useDockerDesktopClient();

    const fetchAndDisplayResponse = async () => {
        const result = await ddClient.extension.vm?.service?.post('/vault', {});
        setResponse(JSON.stringify(result));
    };

    return (
        <>
            <Typography variant="h3">Docker extension demo</Typography>
            <Typography variant="body1" color="text.secondary" sx={{mt: 2}}>
                This
            </Typography>
            <Typography variant="body1" color="text.secondary" sx={{mt: 2}}>
                Pressing the below button will trigger a request to the backend. Its
                response will appear in the textarea.
            </Typography>
            <Stack direction="row" alignItems="start" spacing={2} sx={{mt: 4}}>
                <Button variant="contained" onClick={fetchAndDisplayResponse}>
                    Call backend
                </Button>

                <TextField
                    label="Backend response"
                    sx={{width: 480}}
                    disabled
                    multiline
                    variant="outlined"
                    minRows={5}
                    value={response ?? ''}
                />
            </Stack>
        </>
    );
}
