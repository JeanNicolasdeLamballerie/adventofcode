use egui::{Button, Color32, FontDefinitions, FontFamily, RichText, TextBuffer};

use crate::{
    puzzles::{get_availability, get_puzzle, get_years, Info, INFORMATIONS},
    Part, Steps,
};

#[derive(Default, Clone)]
struct Puzzle {
    auto: bool,
    started: bool,
    ended: bool,
    steps: Steps,
    parts: (bool, bool),
}
#[derive(Default, Clone)]
struct Selection {
    day: Option<usize>,
    year: Option<usize>,
    part: Part,
}
impl Selection {
    fn selected(&self) -> bool {
        return self.day.is_some() && self.year.is_some();
    }
}
#[derive(Clone)]
pub struct VisualizerSettings {
    font_size: usize,
    operations_per_render: usize,
}
impl Default for VisualizerSettings {
    fn default() -> Self {
        Self {
            font_size: 14,
            operations_per_render: 1,
        }
    }
}

#[derive(Default)]
pub struct WindowGui {
    puzzle: Puzzle,
    selection: Selection,
    show_menu: bool,
    expand_info: bool,
    settings: VisualizerSettings,
}

impl WindowGui {
    pub fn new(cc: &eframe::CreationContext<'_>) -> Self {
        let mut fonts = FontDefinitions::default();
        // Install my own font (maybe supporting non-latin characters):
        fonts.font_data.insert(
            "FiraCodeMonoRegular".to_owned(),
            egui::FontData::from_static(include_bytes!("./fonts/FiraCodeMonoRegular.ttf")),
        ); // .ttf and .otf supported

        // // Put my font first (highest priority):
        // fonts
        //     .families
        //     .get_mut(&FontFamily::Proportional)
        //     .unwrap()
        //     .insert(0, "my_font".to_owned());

        // Put my font as last fallback for monospace:
        fonts
            .families
            .get_mut(&FontFamily::Monospace)
            .unwrap()
            .push("FiraCodeMonoRegular".to_owned());

        cc.egui_ctx.set_fonts(fonts);
        // Customize egui here with cc.egui_ctx.set_fonts and cc.egui_ctx.set_visuals.
        // Restore app state using cc.storage (requires the "persistence" feature).
        // Use the cc.gl (a glow::Context) to create graphics shaders and buffers that you can use
        // for e.g. egui::PaintCallback.
        Self::default()
    }
    fn reset(&mut self) {
        self.puzzle = Puzzle::default();
    }
}

impl eframe::App for WindowGui {
    fn update(&mut self, ctx: &egui::Context, _frame: &mut eframe::Frame) {
        egui::TopBottomPanel::top("main top panel").show(ctx, |ui| {
            ui.vertical_centered(|ui| {
                ui.label(RichText::new("Welcome to Advent of Code !").size(30f32));

                if self.show_menu {
                    ui.separator();
                    for available_year in get_years() {
                        let mut button_text = RichText::new(format!("Year : {}", available_year));
                        if let Some(selected_year) = self.selection.year {
                            if selected_year == *available_year {
                                button_text =
                                    button_text.color(Color32::from_hex("#00FFFF").unwrap());
                            }
                        }
                        if ui.button(button_text.heading().size(25f32)).clicked() {
                            self.selection.year = Some(*available_year);
                        }
                    }

                    let mut layout = egui::Layout::left_to_right(egui::Align::TOP);
                    layout = layout.with_main_wrap(true);
                    layout.main_align = egui::Align::Center;

                    ui.separator();
                    ui.with_layout(layout, |ui| {
                        if let Some(year) = self.selection.year {
                            for (day, availability) in get_availability(year).iter().enumerate() {
                                let mut button_text =
                                    RichText::new(format!("Day {}", day + 1)).size(19f32);
                                if let Some(selected_day) = self.selection.day {
                                    if day == selected_day {
                                        button_text = button_text
                                            .color(Color32::from_hex("#00FFFF").unwrap());
                                    }
                                }

                                if ui
                                    .add_enabled(
                                        availability.0 || availability.1,
                                        Button::new(button_text),
                                    )
                                    .clicked()
                                {
                                    self.selection.day = Some(day);
                                    self.puzzle.parts = *availability;
                                }
                            }
                        }
                    });
                };
                let message = if self.expand_info {
                    "Hide Infos"
                } else {
                    "Show Infos"
                };
                if ui.button(message).clicked() {
                    self.expand_info = !self.expand_info;
                }
                let button_icon = if self.show_menu { "^" } else { "v" };
                if ui.button(button_icon).clicked {
                    self.show_menu = !self.show_menu;
                };
            });
        });
        let target = if self.selection.day.is_some() && self.selection.year.is_some() {
            format!(
                "{}-{}",
                self.selection.year.unwrap(),
                self.selection.day.unwrap()
            )
        } else if self.selection.year.is_some() {
            self.selection.year.unwrap().to_string()
        } else {
            "".into()
        };
        let default: &[Info] = &DEFAULT_NOTES;
        let info_list = INFORMATIONS.get(target.as_str()).unwrap_or(&default);
        egui::CentralPanel::default().show(ctx, |ui| {
            if self.selection.selected() {
                egui::TopBottomPanel::top("settings").show_animated(ctx, true, |ui| {
                    ui.horizontal_wrapped(|ui| {
                        ui.vertical(|ui| {
                            ui.add(
                                egui::Slider::new(
                                    &mut self.settings.operations_per_render,
                                    1..=5000,
                                )
                                .logarithmic(true)
                                .text("Operations per Step (Animation speed)"),
                            );
                            ui.add(
                                egui::Slider::new(&mut self.settings.font_size, 3..=100)
                                    .logarithmic(false)
                                    .text("Text size"),
                            );
                        });
                        ui.separator();
                        let start_stop_button = if self.puzzle.started && self.puzzle.auto {
                            "stop"
                        } else {
                            "start"
                        };

                        let mut part_1 = RichText::new("Solve Part 1").heading();
                        let mut part_2 = RichText::new("Solve Part 2").heading();
                        match self.selection.part {
                            Part::One => {
                                part_1 = part_1.color(Color32::from_hex("#00FFFF").unwrap())
                            }
                            Part::Two => {
                                part_2 = part_2.color(Color32::from_hex("#00FFFF").unwrap())
                            }
                            Part::None => {}
                        }

                        if ui
                            .add_enabled(self.puzzle.parts.0, egui::Button::new(part_1))
                            .clicked()
                        {
                            self.selection.part = Part::One;

                            let mut puzzle = get_puzzle(
                                self.selection.year.unwrap(),
                                self.selection.day.unwrap(),
                            );
                            puzzle.process();
                            let solved = puzzle.solve_p1();
                        };
                        if ui
                            .add_enabled(self.puzzle.parts.1, egui::Button::new(part_2))
                            .clicked()
                        {
                            self.selection.part = Part::Two;
                            let puzzle = get_puzzle(
                                self.selection.year.unwrap(),
                                self.selection.day.unwrap(),
                            );
                        };
                        ui.separator();
                        if ui
                            .add(egui::Button::new(
                                RichText::new(start_stop_button).heading(),
                            ))
                            .clicked()
                        {
                            self.puzzle.auto = !self.puzzle.auto;
                        };
                    });
                });
            };

            egui::SidePanel::right("Info Panel").show_animated(ctx, self.expand_info, |ui| {
                for info in info_list.iter() {
                    ui.label(info.message);
                }
            });
        });
    }
}
const DEFAULT_NOTES: [Info; 1] = [Info {
    message: "No notes...",
}];
