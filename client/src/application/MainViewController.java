package application;

import java.io.File;

import javax.imageio.ImageIO;

import javafx.embed.swing.SwingFXUtils;
import javafx.fxml.FXML;
import javafx.scene.canvas.Canvas;
import javafx.scene.canvas.GraphicsContext;
import javafx.scene.control.CheckBox;
import javafx.scene.control.TextField;
import javafx.scene.image.Image;
import javafx.scene.control.ColorPicker;

public class MainViewController {
	
	@FXML
	private Canvas canvas;
	
	@FXML
	private ColorPicker ColorPicker;
	
	@FXML
	private TextField brushSize;
	
	@FXML
	private CheckBox eraser;
	
	
	public void initialize() {
		GraphicsContext g = canvas.getGraphicsContext2D();
		canvas.setOnMouseDragged(e ->{
			double size = Double.parseDouble(brushSize.getText());
			double x = e.getX() - size / 2;
			double y = e.getY() - size / 2;
			
			if(eraser.isSelected()) {
				g.clearRect(x, y, size, size);
			} else {
				g.setFill(ColorPicker.getValue());
				g.fillRect(x, y, size, size);
			}
		});
	}
	

}
