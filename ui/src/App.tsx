import React from 'react';
import Button from '@mui/material/Button';
import {createDockerDesktopClient} from '@docker/extension-api-client';
import {Divider, Grid, TextField, Typography} from '@mui/material';

const client = createDockerDesktopClient();

function useDockerDesktopClient() {
    return client;
}

export function App() {
    const [response, setResponse] = React.useState<string>();
    const [unlockPass, setUnlockPass] = React.useState<string>();
    const [vaultServer, setVaultServer] = React.useState<string>("http://host.docker.internal:8080");
    const [folderID, setFolderID] = React.useState<string>();
    const ddClient = useDockerDesktopClient();

    const fetchAndDisplayResponse = async () => {
        try {
            const result = await ddClient.extension.vm?.service?.post('/vault', {
                "unlock": unlockPass,
                "url": vaultServer,
                "folder_id": folderID,
            });
            setResponse(JSON.stringify(result));
        } catch ({message}) {
            setResponse(JSON.parse(JSON.stringify(message)));
        }
    };
    return (
        <>
            <Typography variant="h3">Bitwarden Vault Extension</Typography>
            <Typography variant="body1" color="text.primary" sx={{mt: 2}}>
                This is a Vault example using Bitwarden. To start set the vault password.
                You need to follow these steps:
                <ol>
                    <li>You need to install the <a href="https://bitwarden.com/help/cli/" target="_blank">Bitwarden
                        CLI.</a></li>
                    <li>Login to cli using <code>bw login</code></li>
                    <li>Finally, run the internal bitwarden server <code>bw serve --port 8080 --hostname all</code></li>
                </ol>
            </Typography>
            <Divider/>
            <Grid container spacing={2}>
                <Grid item xs={4}>
                    <TextField
                        fullWidth
                        id="password"
                        label="Vault Server"
                        variant="outlined"
                        type="text"
                        defaultValue={vaultServer}
                        value={vaultServer}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setVaultServer(event.target.value);
                        }}
                    />
                </Grid>
                <Grid item xs={3}>
                    <TextField
                        fullWidth
                        id="folderID"
                        label="Folder ID"
                        variant="outlined"
                        type="text"
                        defaultValue={folderID}
                        value={folderID}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setFolderID(event.target.value);
                        }}
                    />
                </Grid>
                <Grid item xs={3}>
                    <TextField
                        fullWidth
                        id="password"
                        label="Vault Password"
                        variant="outlined"
                        type="password"
                        defaultValue=""
                        value={unlockPass}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setUnlockPass(event.target.value);
                        }}
                    />
                </Grid>
                <Grid item xs={2}>
                    <Button
                        fullWidth
                        variant="contained"
                        onClick={fetchAndDisplayResponse}>
                        Generate
                    </Button>
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        label="Backend response"
                        disabled
                        multiline
                        variant="outlined"
                        minRows={5}
                        value={response ?? ''}
                    />
                </Grid>
            </Grid>
        </>
    );
}
