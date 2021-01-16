mod data;

pub use crate::data::intake::{Collection, Song};
use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}

#[post("/echo")]
async fn echo(req_body: String) -> impl Responder {
    HttpResponse::Ok().body(req_body)
}

async fn manual_hello() -> impl Responder {
    HttpResponse::Ok().body("Hey there!")
}

#[derive(Serialize, Deserialize, Debug)]
struct Point {
    x: i32,
    y: i32,
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let _point = Point { x: 1, y: 2 };
    let mut v: Vec<Song> = Vec::new();
    v.push(Song { title: "foo".to_string(), song_type: None});
    let col = Collection { 
        title: "Thing".to_string(),
        songs: v,
    };

    let s: String = "{\"title\":\"Thing\",\"songs\":[{\"title\":\"foo\"}]}".to_string();
    
    let serialized = serde_json::to_string(&col).unwrap();
    println!("serialized = {}", serialized);

    let deserialized: Collection = serde_json::from_str(&s).unwrap();
    println!("deserialized = {:?}", deserialized);

    HttpServer::new(|| {
        App::new()
            .service(hello)
            .service(echo)
            .route("/hey", web::get().to(manual_hello))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
