var hiji_kakudo_slider = 0;
var ude_furi_slider = 0;
var hiza_mage_slider = 0;
var asi_furi_slider = 0;
var koshi_kakudo_slider = 0;
var speed = 0;
var my_canvas;
var ude_upleft_angle = 0.0;
var ude_upright_angle = 0.0;
var asi_upleft_angle = 0.0;
var asi_upright_angle = 0.0;
var asi_downleft_angle = 0.0;
var asi_downright_angle = 0.0;
var ude_up_add_angle = 0.0;
var asi_up_add_angle = 0.0;
var asi_down_add_angle = 0.0;
var count = 0.0;
var changing = true;
var i = 0.0;

function setup() {
  frameRate(30);
  my_canvas = createCanvas(800, 400);
  my_canvas.parent("p5container");
  textSize(32);
  fill(0);
  noStroke();
}

function draw() {
  update_slider();
  background(250);
  translate(width / 2, height / 2);
  display_whole_body();
  update_whole_body();
}

function init() {
  count = 0.0;
  ude_up_add_angle = ude_furi_slider * speed / 90.0;
  asi_up_add_angle = asi_furi_slider * speed / 90.0;
  asi_down_add_angle = hiza_mage_slider * speed / 60.0;
  ude_upleft_angle = 0.0;
  ude_upright_angle = 0.0;
  asi_upleft_angle = 0.0;
  asi_upright_angle = 0.0;
  if(changing===true) {
    asi_downleft_angle  = 0.0;
    asi_downright_angle = 0.0;
    for (i=0; i<=30; i+=speed/2.0)
      asi_downright_angle -= asi_down_add_angle;
  }
}

function update_slider() {
  if (did_slider_change()) {
    hiji_kakudo_slider = $('#hiji-kakudo').val();
    ude_furi_slider = $('#ude-furi').val();
    hiza_mage_slider = $('#hiza-mage').val();
    asi_furi_slider = $('#asi-furi').val();
    koshi_kakudo_slider = $('#koshi-kakudo').val();
    speed = $("input[name='sokudo']:checked").val();
    changing = true;
    init();
  } else {
  	if (changing===true && count>30) {
  		changing = false;
  	}
  }
}

function did_slider_change() {
  if (hiji_kakudo_slider != $('#hiji-kakudo').val()) return true;
  if (ude_furi_slider != $('#ude-furi').val()) return true;
  if (hiza_mage_slider != $('#hiza-mage').val()) return true;
  if (asi_furi_slider != $('#asi-furi').val()) return true;
  if (koshi_kakudo_slider != $('#koshi-kakudo').val()) return true;
  if (speed != $("input[name='sokudo']:checked").val()) return true;
  return false;
}

function update_whole_body() {
  counter();
  update_ude_up();
  update_asi_up();
  update_asi_down();
}

function display_whole_body() {
  display_asi_right();
  display_ude_right();
  display_head();
  display_hara();
  display_kosi();
  display_mune();
  display_ude_left();
  display_asi_left();
}

function counter() {
  count += speed/2.0;
  if (count > 120) {
    init();
  }
}

function display_hara() {
  ellipse(0, 0, 25, 25);
}

function display_kosi() {
  beginShape();
  vertex(0, -10);
  vertex(7, 0);
  vertex(8, 35);
  vertex(-10, 35);
  vertex(-7, -10);
  endShape(CLOSE);
}

function display_mune() {
  push();
  rotate(radians(-koshi_kakudo_slider));
  beginShape();
  vertex(0, -7);
  vertex(10, 0);
  vertex(13, -35);
  vertex(-15, -35);
  vertex(-10, -7);
  endShape(CLOSE);
  pop();
}

function display_head() {
  push();
  stroke(0);
  strokeWeight(10);
  rotate(radians(-koshi_kakudo_slider));
  translate(0, -35);
  line(0, 0, 0, -25);
  translate(0, -25);
  ellipse(0, 0, 30, 30);
  pop();
}

function display_ude_left() {
  push();
  stroke(127);
  strokeWeight(4);
  strokeJoin(ROUND);
  rotate(radians(-koshi_kakudo_slider));
  // ude_upleft begin
  translate(11, -33);
  rotate(radians(ude_upleft_angle));
  line(0, 0, 0, 40);
  // ude_upleft end
  // ude_downleft begin
  translate(0, 40);
  rotate(radians(hiji_kakudo_slider));
  line(0, 0, 0, 30);
  translate(0, 30);
  fill(127);
  triangle(0, 0, -3, 10, 5, 8);
  // ude_downleft end

  pop();
}

function display_ude_right() {
  push();
  stroke(127);
  strokeWeight(4);
  strokeJoin(ROUND);
  rotate(radians(-koshi_kakudo_slider));
  // ude_upright begin
  translate(-13, -33);
  rotate(radians(ude_upright_angle));
  line(0, 0, 0, 40);
  // ude_upright end
  // ude_downright begin
  translate(0, 40);
  rotate(radians(hiji_kakudo_slider));
  line(0, 0, 0, 30);
  translate(0, 30);
  fill(127);
  triangle(0, 0, -3, 10, 5, 8);
  // ude_downright end
  pop();
}

function display_asi_left() {
  push();
  stroke(127);
  strokeWeight(4);
  strokeJoin(ROUND);
  // asi_upleft begin
  translate(6, 35);
  rotate(radians(asi_upleft_angle));
  line(0, 0, 0, 40);
  // asi_upleft end
  // asi_downleft begin
  translate(0, 40);
  rotate(radians(asi_downleft_angle));
  line(0, 0, 0, 40);
  translate(0, 40);
  fill(127);
  triangle(0, 0, -10, 10, 5, 10);
  // asi_downleft end
  pop();
}

function display_asi_right() {
  push();
  stroke(127);
  strokeWeight(4);
  strokeJoin(ROUND);
  // asi_upright begin
  translate(-8, 35);
  rotate(radians(asi_upright_angle));
  line(0, 0, 0, 40);
  // asi_upright end
  // asi_downright begin
  translate(0, 40);
  rotate(radians(asi_downright_angle));
  line(0, 0, 0, 40);
  line(0, 0, 0, 40);
  translate(0, 40);
  fill(127);
  triangle(0, 0, -10, 10, 5, 10);
  // asi_downright end
  pop();
}

function update_ude_up() {
  if (count <= 30) {
    ude_upleft_angle += ude_up_add_angle;
    ude_upright_angle -= ude_up_add_angle;
  } else if (count <= 60) {
    ude_upleft_angle -= ude_up_add_angle;
    ude_upright_angle += ude_up_add_angle;
  } else if (count <= 90) {
    ude_upleft_angle -= ude_up_add_angle;
    ude_upright_angle += ude_up_add_angle;
  } else if (count <= 120) {
    ude_upleft_angle += ude_up_add_angle;
    ude_upright_angle -= ude_up_add_angle;
  }
}

function update_asi_up() {
  if (count <= 30) {
    asi_upleft_angle  -= asi_up_add_angle;
    asi_upright_angle += asi_up_add_angle;
  } else if (count <= 60) {
    asi_upleft_angle  += asi_up_add_angle;
    asi_upright_angle -= asi_up_add_angle;
  } else if (count <= 90) {
    asi_upleft_angle  += asi_up_add_angle;
    asi_upright_angle -= asi_up_add_angle;
  } else if (count <= 120) {
    asi_upleft_angle  -= asi_up_add_angle;
    asi_upright_angle += asi_up_add_angle;
  }
}

function update_asi_down() {
  if (count <= 30) {
    asi_downleft_angle = 0.0;
    console.log(asi_downright_angle);
  } else if (count <= 60) {
    asi_downleft_angle -= asi_down_add_angle;
    asi_downright_angle += asi_down_add_angle;
  } else if (count <= 90) {
    asi_downright_angle = 0.0;
  } else if (count <= 120) {
    asi_downleft_angle += asi_down_add_angle;
    asi_downright_angle -= asi_down_add_angle;
  }
}
