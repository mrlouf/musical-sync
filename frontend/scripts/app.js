function checkBackendHealth() {
    fetch('/api/health')
        .then(response => response.text())
        .then(html => {
            document.getElementById('backend-status').innerHTML = html;
        })
        .catch(error => {
            console.error('Error:', error);
            document.getElementById('backend-status').innerHTML = '❌ Error';
        });
}

function getRandomTrack() {
    fetch('/api/track/random')
        .then(response => response.text())
        .then(trackinfo => {
            const container = document.getElementById('status-container');

            const trackData = JSON.parse(trackinfo);

            container.innerHTML = `<div class="sync-status">
                <h2 style="margin-bottom: 20px; color: #333;">Track retrieved:</h2>
                <div class="status-item">
                    <span class="status-label">${trackData.artist}</span>
                    <span class="status-label">${trackData.track}</span>
                    <span class="status-label">${trackData.album}</span>
                </div>
            </div>`;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function getAlbum() {
    fetch('/api/album/random')
        .then(response => response.text())
        .then(albuminfo => {
            const container = document.getElementById('status-container');

            const trackData = JSON.parse(albuminfo);

            console.log(trackData);

            const tracks = trackData.tracks;

            const tracklist = Array.isArray(tracks) ? tracks.map((element, idx) => {
                const title = typeof element === 'string' ? element : (element.title || element.name || JSON.stringify(element));
                return `<div class="status-item">
                            <span class="status-label">${idx + 1}. ${title}</span>
                        </div>`;
            }) : [];

            container.innerHTML = `<div class="sync-status">
                <h2 style="margin-bottom: 20px; color: #333;">Album retrieved:</h2>
                <div class="status-item">
                    <span class="status-label">${trackData.artist}</span>
                    <span class="status-label">${trackData.album}</span>
                </div>
                <div style="margin-top: 15px;">
                    ${tracklist.join('')}
                </div>
            </div>`;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

async function getTrackNumberFromPlaylists() {
    const container = document.getElementById('status-container');
    container.innerHTML = `<div class="sync-status">
        <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
        <div class="status-item">
            <span class="status-label">Loading…</span>
            <span class="status-value">⌛</span>
        </div>
    </div>`;

    try {
        const res = await fetch('/api/playlists/tracknumbers');
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        const data = await res.json();

        const nbDeezer = data && data.nb_tracks_Deezer != null ? data.nb_tracks_Deezer : 'N/A';
        const nbSpotify = data && data.nb_tracks_Spotify != null ? data.nb_tracks_Spotify : 'N/A';

        const content = `<div class="sync-status">
            <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
            <div class="status-item">
                <span class="status-label">Deezer tracks: </span>
                <span class="status-value" id="tracks-deezer">${nbDeezer}</span>
            </div>
            <div class="status-item">
                <span class="status-label">Spotify tracks: </span>
                <span class="status-value" id="tracks-spotify">${nbSpotify}</span>
            </div>
        </div>`;

        container.innerHTML = content;
    } catch (err) {
        console.error('Error fetching playlist:', err);
        container.innerHTML = `<div class="sync-status">
            <div class="status-item">
                <span class="status-label">Error:</span>
                <span class="status-value">Failed to load playlist</span>
            </div>
        </div>`;
    }
}

function loginSpotify() {
    fetch('/api/login/spotify')
        .then(response => response.text())
        .then(html => {
            document.getElementById('status-container').innerHTML = html;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

async function getSyncStatus() {
    const container = document.getElementById('status-container');
    container.innerHTML = `<div class="sync-status">
        <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
        <div class="status-item">
            <span class="status-label">Loading…</span>
            <span class="status-value">⌛</span>
        </div>
    </div>`;

    try {
        const res = await fetch('/api/sync-status');
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        const html = await res.text();

        const content = `<div class="sync-status">
            <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
            <div class="status-item">
                <span class="status-label">All tracks are synchronised: </span>
                <span class="status-value" id="sync-ok">OK</span>
            </div>
        </div>`;
        container.innerHTML = content;
    } catch (err) {
        console.error('Error fetching sync status:', err);
        container.innerHTML = `<div class="sync-status">
            <div class="status-item">
                <span class="status-label">Error:</span>
                <span class="status-value">Failed to load sync status</span>
            </div>
        </div>`;
    }
}

async function getPlaylistFromDeezer() {
    const container = document.getElementById('status-container');
    container.innerHTML = `<div class="sync-status">
        <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
        <div class="status-item">
            <span class="status-label">Loading…</span>
            <span class="status-value">⌛</span>
        </div>
    </div>`;

    try {
        const response = await fetch('/api/playlist/deezer');
        const playlistinfo = await response.text();

        const trackData = JSON.parse(playlistinfo);

        const tracks = trackData.tracks;

        const tracklist = tracks.map((element, idx) => {
            const title = typeof element === 'string' ? element : (element.title || element.name || JSON.stringify(element));
            return `<div class="status-item">
                        <span class="status-label">${idx + 1}. ${title}</span>
                    </div>`;
        });

        container.innerHTML = `<div class="sync-status">
            <h2 style="margin-bottom: 20px; color: #333;">Deezer Playlist retrieved:</h2>
            <div style="margin-top: 15px;">
                ${tracklist.join('')}
            </div>
        </div>`;

    } catch (err) {
        console.error('Error fetching sync status:', err);
        container.innerHTML = `<div class="sync-status">
            <div class="status-item">
                <span class="status-label">Error:</span>
                <span class="status-value">Failed to load sync status</span>
            </div>
        </div>`;
    }
}

async function getPlaylistFromSpotify() {
    const container = document.getElementById('status-container');
    container.innerHTML = `<div class="sync-status">
        <h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
        <div class="status-item">
            <span class="status-label">Loading…</span>
            <span class="status-value">⌛</span>
        </div>
    </div>`;

    try {
        const response = await fetch('/api/playlist/spotify');
        const playlistinfo = await response.text();

        const trackData = JSON.parse(playlistinfo);
        const tracks = trackData.tracks;

        const tracklist = tracks.map((element, idx) => {
            const title = typeof element === 'string' ? element : (element.track?.title || element.track?.name || JSON.stringify(element));
            return `<div class="status-item">
                        <span class="status-label">${idx + 1}. ${title}</span>
                    </div>`;
        });

        container.innerHTML = `<div class="sync-status">
            <h2 style="margin-bottom: 20px; color: #333;">Spotify Playlist retrieved:</h2>
            <div style="margin-top: 15px;">
                ${tracklist.join('')}
            </div>
        </div>`;

    } catch (err) {
        console.error('Error fetching sync status:', err);
        container.innerHTML = `<div class="sync-status">
            <div class="status-item">
                <span class="status-label">Error:</span>
                <span class="status-value">Failed to load sync status</span>
            </div>
        </div>`;
    }
}

document.addEventListener('DOMContentLoaded', function() {
    checkBackendHealth();
});
