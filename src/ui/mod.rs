use window::WindowGui;

pub mod window;

pub enum UserInterface {
    Window,
    Terminal,
}

impl UserInterface {
    pub fn init() -> Self {
        return UserInterface::Window;
    }
    pub fn render(&self) {
        match self {
            UserInterface::Window => {
                let options = eframe::NativeOptions::default();
                eframe::run_native(
                    "Advent of Code",
                    options,
                    Box::new(|cc| Ok(Box::new(WindowGui::new(cc)))),
                )
                .unwrap();
            }
            UserInterface::Terminal => todo!(),
        }
    }
}
