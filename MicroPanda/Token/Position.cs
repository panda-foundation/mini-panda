namespace MicroPanda.Token;

internal class Position
{
    readonly File File;
	internal readonly int Offset;

	internal Position(File file, int offset)
	{
		File = file;
		Offset = offset;
	}

	public override string ToString()
	{
		var path = File.Name;
		var (line, column) = File.GetLocation(Offset);
		return $"{path}:{line}:{column}";
	}

	internal int GlobalOffset() => File.FileOffset + Offset;
}

internal class File
{
	internal readonly string Name;
	internal int Size;
	internal int FileOffset;
	private readonly List<int> LinesOffset = [];

	internal File(string name, int size)
	{
		Name = name;
		Size = size;
		LinesOffset.Add(0);
	}

	internal (int Line, int Column) GetLocation(int offset)
	{
		var i = 0;
		var j = LinesOffset.Count;
		while (i < j)
		{
			var h = i + (j - i) / 2;
			if (LinesOffset[h] <= offset)
			{
				i = h + 1;
			}
			else
			{
				j = h;
			}
		}
		i--;
		if (i >= 0)
		{
			return (i + 1, offset - LinesOffset[i] + 1);
		}
		return (0, 0);
	}

	internal void AddLine(int offset)
	{
		LinesOffset.Add(offset);
	}

	internal int LinesCount() => LinesOffset.Count;

	internal Position GetPosition(int offset) => new(this, offset);

	internal int GetGlobalOffset(int offset) => FileOffset + offset;
}

internal class FileSet
{
    private readonly List<File> Files = [];
	private int FileOffset;

	internal File AddFile(string name, int size)
	{
		foreach (var file in Files)
		{
			if (file.Name == name)
			{
				throw new InvalidOperationException($"File {name} already added");
			}
		}

        var newFile = new File(name, size)
        {
            FileOffset = FileOffset
        };
        FileOffset += size + 1;
		Files.Add(newFile);
		return newFile;
	}

	internal File? GetFile(int globalOffset)
	{
		foreach (var file in Files)
		{
			if (globalOffset <= file.FileOffset + file.Size)
			{
				return file;
			}
		}
		return null;
	}

	internal Position? GetPosition(int globalOffset)
	{
		var file = GetFile(globalOffset);
		if (file != null)
		{
			return file.GetPosition(globalOffset - file.FileOffset);
		}
		return null;
	}

	internal void UpdateFileSize(string name, int size)
	{
		var found = false;
		foreach (var file in Files)
		{
			if (file.Name == name)
			{
				found = true;
				file.Size = size;
			}
		}
		if (found)
		{
			FileOffset = 0;
			foreach (var file in Files)
			{
				file.FileOffset = FileOffset;
				FileOffset += file.Size + 1;
			}
		}
	}
}