pub mod intake {
    use serde::{Deserialize, Serialize};

    #[derive(Serialize, Deserialize, Debug)]
    pub struct Collection {
        pub title: String,
        pub songs: Vec<Song>,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct Song {
        pub title: String,
        pub song_type: Option<String>,
    }
}
