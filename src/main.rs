use egui::Color32;
use ui::UserInterface;

pub mod puzzles;
pub mod ui;

fn main() {
    let ui = UserInterface::init();
    ui.render();
}

trait SolvePuzzle {
    fn solve_p1(&mut self) -> Solved;
    fn solve_p2(&mut self) -> Solved;
}

trait Process {
    fn process(&mut self);
}

trait Solvable: SolvePuzzle + Process {}
impl<T> Solvable for T where T: SolvePuzzle + Process {}

/// Results for the puzzle.
enum Solved {
    Usize(usize),
    Usteps(usize, Steps),
    I64steps(i64, Steps),
}

#[derive(Clone, Default)]
/// Each individual input character that will be displayed.
struct StepInputChar {
    char: String,
    color: Option<Color32>,
    bg_color: Option<Color32>,
}

#[derive(Clone)]
/// The action that should be taken on the input text, e.g, changing it's color, it's content...
enum TextAction {
    Color(Color32),
    Replace(String),
    ReplaceWithColor(Color32, String),
}

#[derive(Clone)]
/// The text-based action, and the target in the input (`StepInputChar`). Careful when matching,
/// since the target refers to the original input, not one with removed newlines/whitespaces.
struct StepAction {
    target: usize,
    action: TextAction,
}

#[derive(Clone, Default)]
// The steps of the visualization.
struct Steps {
    input: Vec<StepInputChar>,
    actions: Vec<StepAction>,
}

#[derive(Default, Clone)]
enum Part {
    One,
    Two,
    #[default]
    None,
}

pub fn result_to_display<T>(number: T) -> String
where
    T: ToString,
{
    use unicode_segmentation::UnicodeSegmentation;
    let mut count = 0;
    let mut string_displayed = String::new();
    for graph in number.to_string().graphemes(true).rev() {
        if count != 0 && count % 3 == 0 {
            string_displayed.push(' ');
        }
        string_displayed.push_str(graph);
        count += 1;
    }
    string_displayed.graphemes(true).rev().collect()
}

impl From<String> for Steps {
    fn from(value: String) -> Self {
        use unicode_segmentation::UnicodeSegmentation;
        let input = value
            .graphemes(true)
            .map(|grapheme| StepInputChar {
                char: grapheme.to_string(),
                color: None,
                bg_color: None,
            })
            .collect();

        return Steps {
            input,
            actions: Vec::new(),
        };
    }
}
