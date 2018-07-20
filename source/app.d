import std.stdio;
import std.regex;
import std.file;
import std.path;
import std.string;
import std.array;
import std.algorithm;

void replaceInterfaceVersion()
{
	immutable string tocFile = getcwd.baseName ~ ".toc";

	if(tocFile.exists)
	{
		auto lines = tocFile.readText.lineSplitter();
		string[] outputLines;

		foreach(line; lines)
		{
			if(line.canFind("Interface:"))
			{
				auto re = regex(r"(\w+)(\d+)","g");
				immutable string replacedValue = replaceAll(line, re, "80000");

				outputLines ~= replacedValue;
			}
			else
			{
				outputLines ~= line;
			}
		}

		foreach(line; outputLines)
		{
			writeln(line);
		}
	}
}

void main(string[] arguments)
{
	replaceInterfaceVersion();
}
