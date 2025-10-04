# Diary System Documentation

## Current Implementation

### Storage Strategy
- **Sample Entries**: Stored in Go code (`api/diary-entries.go`)
  - 2 hardcoded entries (1 public, 1 private)
  - Returns from `/api/diary-entries` endpoint
  
- **User-Created Entries**: Stored in browser `localStorage`
  - Saved when you click "Save Entry"
  - Merged with API entries when displaying
  - **NOT synced to Redis/database**

### Data Flow
```
1. Page loads → Fetch from /api/diary-entries (sample entries)
2. Merge with localStorage (your entries)
3. Sort by date (newest first)
4. Display all entries
5. Private entries show "[Private]" tag
6. Private content censored unless logged in
```

### Security
- Server-side censoring: Private entries return "This entry is private. Login to view." when not authenticated
- Can't bypass with inspect element (content never sent to browser)
- Password stored in Redis only (`diary_password` key)

### What's NOT Implemented Yet
- ❌ Saving new entries to Redis/database
- ❌ POST endpoint to create entries
- ❌ Syncing entries across devices
- ❌ Entry editing/deletion

### Current State
✅ View sample entries (from API)
✅ Create new entries (localStorage only)
✅ Public/private entry system
✅ Server-side censoring
✅ Password authentication
✅ Clean UI with no emojis
❌ Entries persist only in browser localStorage

## Setup Required

1. Set diary password in Redis:
   ```bash
   # In Vercel Redis Data Browser
   SET diary_password "your_secret_password"
   ```

2. Test the diary:
   - Visit `/diary`
   - See 2 sample entries (1 public, 1 private)
   - Login to view private content
   - Create new entries (saved to localStorage)

## Future Enhancement Options

If you want entries saved to Redis:
1. Create POST endpoint in `api/diary-entries.go`
2. Store entries in Redis as JSON array
3. Update frontend to POST new entries to API
4. Remove localStorage dependency
