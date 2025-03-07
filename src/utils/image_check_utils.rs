use std::{fs, path::Path};

use super::status_msg::error;

pub fn get_image_type() -> String {
    let image_type_file: Option<&Path>;

    if Path::new("/usr/share/horizon").exists() {
        image_type_file = Some(Path::new("/usr/share/horizon/image_type"));
    } else if Path::new("/usr/share/nova/image_type").exists() {
        image_type_file = Some(Path::new("/usr/share/nova/image_type"));
    } else if Path::new("/usr/share/umbra/image_type").exists() {
        image_type_file = Some(Path::new("/usr/share/umbra/image_type"));
    } else {
        // Return an error if no image type file is found
        error("No image type file found!");
        panic!()
    }

    let image = fs::read_to_string(image_type_file.unwrap()).expect("Failed to get image type");

    let image_name = image.trim().to_string();

    return image_name;
}

pub fn is_correct_image() -> bool {
    if Path::new("/usr/share/horizon").exists() {
        return true;
    } else if Path::new("/usr/share/nova/image_type").exists() {
        return true;
    } else if Path::new("/usr/share/umbra/image_type").exists() {
        return true;
    } else {
        return false;
    }
}
