# musical-sync

A backend-oriented project to synchronize music playlists across platforms.

---

## Why this project?

In 2025, I migrated from Spotify to Deezer as a premium user for various reasons - pricing policies and sound quality being the main ones.

I used a publicly available platform to copy all my playlists from Spotify to ensure I had a backup, even though I do not plan on deleting my account on either platform.  
The main issue I encountered was with a collaborative playlist shared with a Spotify user: I made a copy of the playlist, but adding a song on Deezer does not reflect on Spotify, and vice versa.

Some public apps offer cross-platform playlist synchronisation, which is neat, but most features, even basic ones, are locked behind premium plans.  
Also, I was not particularly happy with the performance of the apps I tried.

That’s when I had the idea to develop my own service. After all, unhappy software engineers are responsible for half the tech stack we use today, right?

---

## Tech used

Since I am focusing on backend engineering, I chose **Golang** as the language for the backend server. The syntax is pretty intuitive, especially for someone accustomed to good ol’ C and compiled languages in general. Some aspects - like package imports or nested structs for JSON handling - are a bit more complex, but overall the learning curve is fairly smooth.

The frontend is intentionally minimal: static **HTML + CSS + vanilla JavaScript**. This gets the job done, and I am definitely not going to spend time learning React or Vue for a project of this size.  
I am considering **HTMX** as a lightweight alternative to rely even more on server-side logic, but this is not a top priority.

The application is containerized with **Docker** for portability, with **NGINX** acting as a reverse proxy.

---

## Architecture overview

```text
+---------------------------+
|Frontend (HTML / CSS / JS) |
+---------------------------+
              |
        HTTP Requests
              |
              v
+---------------------------+
|           NGINX           |
|      (Reverse Proxy)      |
+---------------------------+
              |
          REST API
              |
              v
+---------------------------+
|       Go API Server       |
+---------------------------+
              |
        External APIs
              |
              v
+---------------------------+
|   Spotify / Deezer APIs   |
+---------------------------+
```

---

## Current status

The application is able to retrieve information about an artist, an album, or a playlist from both Deezer and Spotify using their respective APIs.

Unfortunately, Deezer has stopped the creation of new developer applications. As a result, the synchronisation is currently **uni-directional**, meaning playlists can only be synchronised from **Deezer to Spotify**.

A reverse synchronization mechanism may be achievable through manual updates, but this is not on the agenda for now.

---

## Future objectives / TODO

- Retrieve the tracklist from a playlist and compare it to another playlist
- Add user authentication via OAuth
- Implement actual playlist synchronisation logic
